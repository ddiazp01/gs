package handlers

import "net/http"

//PathInicio ruta raiz
const PathInicio string = "/"

//PathHome ruta home
const PathHome string = "/home"

//PathJSFiles ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de estilos css
const PathCSSFiles string = "/css/"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathRegister ruta de register
const PathRegister string = "/register"

//PathLoginFile ruta de login
const PathLoginFile string = "/loginFile"

//PathLogin ruta de login
const PathLogin string = "/login"

//PathPerfilFile ruta perfil
const PathPerfilFile string = "/perfil"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//  Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathHome] = HomeFile
	Manejadores[PathJSFiles] = JSFile
	Manejadores[PathCSSFiles] = CSSFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLogin] = Login
	Manejadores[PathPerfilFile] = PerfilFile

}
