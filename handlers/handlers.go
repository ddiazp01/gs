package handlers

import "net/http"

//PathInicio ruta raiz
const PathInicio string = "/"

<<<<<<< HEAD
//PathHome ruta home
const PathHome string = "/home"
=======
//PathIndex ruta raiz
const PathIndex string = "/"
>>>>>>> 354f70492f3b4a8fee63086dff16d623ddaedd6a

//PathJSFiles ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

const pathCSS string = "/css/"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathRegister ruta de register
const PathRegister string = "/register"

//PathLoginFile ruta de login
const PathLoginFile string = "/loginFile"

//PathLogin ruta de login
const PathLogin string = "/login"

//PathDeportesFile ruta perfil
const PathDeportesFile string = "/deportes"

//PathEmpleoFile ruta perfil
const PathEmpleoFile string = "/empleo"

//PathTramitesFile ruta perfil
const PathTramitesFile string = "/tramites"

//PathLogout te lleva al logout
const PathLogout string = "/logout"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

// Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
<<<<<<< HEAD
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathHome] = HomeFile
=======
	Manejadores[PathIndex] = IndexFile
>>>>>>> 354f70492f3b4a8fee63086dff16d623ddaedd6a
	Manejadores[PathJSFiles] = JSFile
	Manejadores[pathCSS] = CSSFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLogin] = Login
	Manejadores[PathTramitesFile] = TramitesFile
	Manejadores[PathDeportesFile] = DeportesFile
	Manejadores[PathEmpleoFile] = EmpleoFile
	Manejadores[PathLogout] = Logout

}
