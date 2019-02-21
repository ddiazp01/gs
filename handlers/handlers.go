package handlers

import "net/http"

//PathIndex Ruta raíz
const PathIndex string = "/"

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
	Manejadores[PathIndex] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CSSFile

}
