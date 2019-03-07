package handlers

import "net/http"

//PathInicio ruta raiz
const PathInicio string = "/"

//PathIndex ruta raiz
const PathIndex string = "/"

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

//PathLogout te lleva al logout
const PathLogout string = "/logout"

<<<<<<< HEAD
//PathInsertCita
const PathInsertCita string = "/insertcita"

//PathListCitas
const PathListCitas string = "/listcitas"

//PathCitasFile
const PathCitasFile string = "/citas"

=======
//PathCitasFile te lleva al logout
const PathCitasFile string = "/citas"

//PathEnvioPeticion2 Ruta de envío de peticiones
const PathEnvioPeticion2 string = "/envio2"

>>>>>>> 8cda52efc8a5d19aabafe99f889d269cfad83798
//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

// Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathIndex] = IndexFile
	Manejadores[PathJSFiles] = JSFile
	Manejadores[pathCSS] = CSSFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLogin] = Login
	Manejadores[PathLogout] = Logout
<<<<<<< HEAD
	Manejadores[PathInsertCita] = InsertCita
	/*Manejadores[PathListCitas] = ListCitas*/
	Manejadores[PathCitasFile] = CitasFile
=======
	Manejadores[PathCitasFile] = CitasFile
	Manejadores[PathEnvioPeticion2] = Insert2

>>>>>>> 8cda52efc8a5d19aabafe99f889d269cfad83798
}
