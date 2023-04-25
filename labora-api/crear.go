package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Crear el enrutador y  luego definir las rutas en la función definirRutas

	enrutador := mux.NewRouter()

	enrutador.HandleFunc("/hola", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("¡Hola, mundo!"))
	}).Methods("GET")
}
