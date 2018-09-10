package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var (

	//ErrNotFound es retornado cuando no existe el indice o el item en la bd
	ErrRequired = errors.New("is required")
)

// Articles struct (Model)
type Articles struct {
	id    string  `json:"id"`
	title string  `json:"title"`
	price float32 `json:"price"`
}

// Car struct
type Carts struct {
	Id    string `json:"id"`
	Owner string `json:"owner"`
}

type CartsArticles struct {
	id       string              `json:"id"`
	articles map[string]Articles `json:"articles"`
}

var arcticles []Articles
var cartsA []Carts
var dbName = "db.json"

// Get all carts
func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cartsA)

}

// Create single cart
func createCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cart Carts
	error := json.NewDecoder(r.Body).Decode(&cart)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(cart.Id) == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(cart.Owner) == "" {
		http.Error(w, "owner is required", http.StatusBadRequest)
		return
	}

	cartsA = append(cartsA, cart)
	json.NewEncoder(w).Encode(cart)
}

// Main function
func main() {
	// Init router

	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/", getCar).Methods("GET")
	r.HandleFunc("/carts/{id}", getCar).Methods("GET")
	r.HandleFunc("/carts", getCar).Methods("GET")
	r.HandleFunc("/carts", createCart).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8003", r))
}
