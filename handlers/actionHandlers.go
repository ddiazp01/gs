package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	client "gs/data/dataclient"
	"gs/data/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	websocket "websocket-master"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (name string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			name = cookieValue["name"]
		}
	}
	return name
}

//NombreUsuario Función que muestra el nombre de usuario logueado
func NombreUsuario(response http.ResponseWriter, request *http.Request) {
	name := getUserName(request)
	fmt.Fprintf(response, name)
}
func setSession(username string, response http.ResponseWriter) {
	value := map[string]string{
		"username": username,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//Login Función para acceder a la página
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLogin {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	respuesta := false
	if e == nil {
		// datos que recibe del cliente
		var usuario model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.UserName)

		if usuario.UserName == "" || usuario.Password == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Contraseña de la base de datos
		password := client.LogearUsuario(&usuario)

		// Comprueba que las dos contraseñas sean iguales
		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(usuario.Password)); err != nil {
			fmt.Printf("No has podido inicar sesión")
		} else {
			respuesta = true
			setSession(usuario.UserName, w)
			fmt.Println("Inicio de sesión realizado")
			getUserName(r)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, respuesta)
	}

	fmt.Fprintln(w, respuesta)
}

//Insert Funcion inserta en la base de datos
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPeticion {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var usuario model.Usuario
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.UserName)

		if usuario.Name == "" || usuario.Apellidos == "" || usuario.UserName == "" || usuario.Password == "" || usuario.Email == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//para incriptar contraseña
		hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		usuario.Password = hashComoCadena
		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(usuario)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPeticion(&usuario)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}

}

const (
	writeWait = 10 * time.Second
	//Tiempo permitido para leer el siguiente mensaje pong del par.
	pongWait = 60 * time.Second
	// Enviar pings para mirar con este período. Debe ser menos de pong Espera.
	pingPeriod = (pongWait * 9) / 10
	// Tamaño máximo de mensaje permitido desde el par.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//El cliente es un intermediario entre la conexión websocket y el hub.
type Usuario struct {
	Name      string
	Apellidos string
	UserName  string
	Password  string
	Email     string
	hub       *Hub
	conn      *websocket.Conn
	send      chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.

func (c *Usuario) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// Escriba los mensajes de Pump pump desde el concentrador a la conexión websocket.
// Se inicia un goroutine que ejecuta writePump para cada conexión. los
// la aplicación garantiza que hay como máximo un escritor a una conexión por
// ejecutando todas las escrituras de este goroutine.
func (c *Usuario) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Servir W maneja las solicitudes de websocket desde el par.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	usuario := &Usuario{hub: hub, conn: conn, send: make(chan []byte, 256)}
	usuario.hub.register <- usuario

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go usuario.writePump()
	go usuario.readPump()
}

// Hub struct
type Hub struct {
	// Registered clients.
	usuarios map[*Usuario]bool
	// Inbound messages from the clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *Usuario
	// Unregister requests from clients.
	unregister chan *Usuario
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Usuario),
		unregister: make(chan *Usuario),
		usuarios:   make(map[*Usuario]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case usuario := <-h.register:
			h.usuarios[usuario] = true
		case usuario := <-h.unregister:
			if _, ok := h.usuarios[usuario]; ok {
				delete(h.usuarios, usuario)
				close(usuario.send)
			}
		case message := <-h.broadcast:
			for usuario := range h.usuarios {
				select {
				case usuario.send <- message:
				default:
					close(usuario.send)
					delete(h.usuarios, usuario)
				}
			}
		}
	}
}

//Logout funcion de salir
func Logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}
