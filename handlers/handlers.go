package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathRegister te lleva al register
const PathRegister string = "/register"

//PathInsert te lleva a la funcion al register
const PathInsert string = "/insert"

//PathLoginFile te lleva al login
const PathLoginFile string = "/loginFile"

//PathLogin te lleva al login
const PathLogin string = "/login"

//PathFotoFile lleva al foto
const PathFotoFile string = "/foto"

//PathListarFoto lleva al foto
const PathListarFoto string = "/listarfoto"


//PathUploader te sube la foto
const PathUploader string = "/uploader"

//PathLogout te lleva al logout
const PathLogout string = "/logout"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de link de css
const PathCSSFiles string = "/css/"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathFotoFile] = FotoIndex
	Manejadores[PathInsert] = Insert
	Manejadores[PathListarFoto] = ListarFoto
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLogin] = Login
	Manejadores[PathLogout] = Logout
	Manejadores[PathUploader] = InsertUploader
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CSSFile


}
