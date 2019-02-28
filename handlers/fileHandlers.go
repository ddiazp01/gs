package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

//IndexFile Función que devuelve el index.html
func IndexFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathIndex {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/index.html")
}

// HomeFile pagina de chat
func HomeFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathHome {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/home.html")
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

//CSSFile Manejador de archivos CSS
func CSSFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	file := r.URL.Path

	if strings.HasPrefix(file, "/") {
		file = file[1:len(r.URL.Path)]
	}

<<<<<<< HEAD
	switch file {
	case //Internos
		"css/register.css",
		"css/base.css":
=======
	file := r.URL.Path

	if strings.HasPrefix(file, "/") {
		file = file[1:len(r.URL.Path)]
	}

	switch file {
	case //Internos
		"css/base.css",
		"css/register.css",
		"css/login.css",
		"css/deportes.css",
		"css/empleo.css",
		"css/tramites.css",
		"css/inicio.css":
>>>>>>> 354f70492f3b4a8fee63086dff16d623ddaedd6a
		http.ServeFile(w, r, file)
		break
	default:
		http.NotFound(w, r)
		return
	}
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

//DeportesFile cargar pagina de parfil
func DeportesFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathDeportesFile {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/deportes.html")
}

//EmpleoFile cargar pagina de parfil
func EmpleoFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathEmpleoFile {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/empleo.html")
}

//TramitesFile cargar pagina de parfil
func TramitesFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud entrante de " + r.URL.EscapedPath())
	if r.URL.Path != PathTramitesFile {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/tramites.html")
}
