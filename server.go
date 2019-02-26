package main

import (
	"fmt"
	hnd "gs/handlers"
	"log"
	"net/http"
	"strconv"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
func main() {
	port := 8080

	fmt.Println("holi")

	for path, handler := range hnd.Manejadores {
		http.HandleFunc(path, handler)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	fmt.Println("Servidor abierto en http://localhost:" + strconv.Itoa(port))
}
