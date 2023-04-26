package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "¡Hola, mundooooo!")
	// })

	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")

	//Más adelante.
	// router.HandleFunc("/items", createItem).Methods("POST")
	// router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	// router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Paula"},
	{ID: "2", Name: "Lucas"},
	{ID: "3", Name: "Misa"},
	{ID: "4", Name: "Rosario"},
	{ID: "5", Name: "Epik High"},
	{ID: "6", Name: "Pepe"},
	{ID: "7", Name: "Misa2"},
	{ID: "8", Name: "Rosario siempre estuvo cerca"},
	{ID: "9", Name: "Bokita"},
	{ID: "10", Name: "Burzaco"},
}

func getItems(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	v := query.Get("id")

	for _, item := range items {
		if item.ID == v {
			json.NewEncoder(w).Encode(item.Name)
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}
