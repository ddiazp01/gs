package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

//IndexFile Función que devuelve el index.html
func IndexFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathInicio {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/index.html")
}

//JSFile Manejador de archivos javascript
func JSFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	file := r.URL.Path

	if strings.HasPrefix(file, "/") {
		file = file[1:len(r.URL.Path)]
	}

	switch file {
	//Externos
	case "js/libs/jquery-3.3.1.min.js",
		"js/libs/moment.min.js",
		//Internos
		"js/base.js":
		http.ServeFile(w, r, file)
		break
	default:
		http.NotFound(w, r)
		return
	}

}

//CSSFile Manejador de archivos Css
func CSSFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "css/base.css")
}

//RegisterFile para abrir pagina de registro
func RegisterFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathRegister {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/register.html")
}

// LoginFile para abrir pagina de login
func LoginFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathLoginFile {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/login.html")
}

//InicioFile cargar pagina de parfil
func InicioFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathInicioFile {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/inicio.html")
}
