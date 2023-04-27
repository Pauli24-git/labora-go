package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	//router.HandleFunc("/items/{id}", getItem)
	router.HandleFunc("/item", getItem).Methods("GET")

	//Más adelante.
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/", updateItem).Methods("PUT")
	router.HandleFunc("/items/", deleteItem).Methods("DELETE")

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
	{ID: "6", Name: "Paula"},
	{ID: "7", Name: "Misa"},
	{ID: "8", Name: "Rosario siempre estuvo cerca"},
	{ID: "9", Name: "Bokita"},
	{ID: "10", Name: "Burzaco"},
}

// func getItems(w http.ResponseWriter, r *http.Request) {

// 	json.NewEncoder(w).Encode(items)
// }

func getItems(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := request.URL.Query() // page - itemsPerPage

	// page := params["page"]
	// itemsPerPage := params["itemsPerPage"]

	page := params.Get("page")
	itemsPerPage := params.Get("itemsPerPage")

	if page == "0" {
		page = "1"
	}
	if itemsPerPage == "0" {
		itemsPerPage = "3"
	}

	pageIndex, _ := strconv.Atoi(page)
	itemsPerPageInt, _ := strconv.Atoi(itemsPerPage)

	var newListItems []Item

	init := itemsPerPageInt * (pageIndex - 1)
	limit := init + itemsPerPageInt

	nroPage := float64(len(items)) / float64(itemsPerPageInt)
	nroPage = math.Ceil(nroPage)

	if pageIndex <= int(nroPage) {
		if limit > len(items) {
			newListItems = items[init:]
		} else {
			newListItems = items[init:limit]
		}
	}
	// Función para obtener todos los elementos
	json.NewEncoder(response).Encode(newListItems)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	v := query.Get("id")
	n := query.Get("name")

	encontrado := false

	for _, item := range items {
		if (item.ID) == v || strings.ToLower(item.Name) == strings.ToLower(n) {
			encontrado = true
			json.NewEncoder(w).Encode(item)
		}
	}

	if !encontrado {
		json.NewEncoder(w).Encode("No se encontro ningun registro con ese id o nombre")
	}
	// params := mux.Vars(r)
	// for _, item := range items {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item.Name)
	// 	}
	// 	if item.ID == params["name"] {
	// 		json.NewEncoder(w).Encode(item.Name)
	// 	}
	// }
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var decodedItem Item
	err := json.NewDecoder(r.Body).Decode(&decodedItem)

	//no funciona, consultar despues porque no captura error o porque deja sumar cosas vacias
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic("panikeamo")
	}

	items = append(items, decodedItem)
	json.NewEncoder(w).Encode(items)

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}
