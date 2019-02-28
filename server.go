package main

import (
	"fmt"
	hnd "gs/handlers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080

	fmt.Println("Servidor de gestion de servicios iniciado")

	for path, handler := range hnd.Manejadores {
		http.HandleFunc(path, handler)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	fmt.Println("Servidor abierto en http://localhost:" + strconv.Itoa(port))

}
