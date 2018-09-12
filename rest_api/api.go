package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jcasmer/GoBootcamp/memory_db/db"
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
	Id      string      `json:"id"`
	Owner   string      `json:"owner"`
	Article []*Articles `json:"articles"`
}

type CartsArticles struct {
	id       string              `json:"id"`
	articles map[string]Articles `json:"articles"`
}

var arcticles []Articles
var cartsA []Carts
var dbName = "db.json"

// Get all carts
func getCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cartsA)

}

// Create single cart
func createCart(w http.ResponseWriter, r *http.Request) {
	// create a sinle cart without articles
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

	// cart.Id = strconv.Itoa(rand.Intn(100000000))
	// cartsA = append(cartsA, cart)

	value, _ := json.Marshal(cart)

	d, erro := db.OpenDB(dbName)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}
	resul := d.CreateWithIndex(cart.Id, string(value))
	if resul != nil {
		http.Error(w, resul.Error(), http.StatusBadRequest)
		return
	}

	er := d.Close(dbName)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
	// http. .StatusCreated

}

func getCart(w http.ResponseWriter, r *http.Request) {
	// retrive a specific cart with its articles
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	d, erro := db.OpenDB(dbName)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}
	var cart Carts
	car, erro := d.Retrieve(params["id"])
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusNotFound)
		return
	}

	err := json.Unmarshal([]byte(car), &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cart)

}

// Main function
func main() {
	// Init router

	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/carts/{id}", getCart).Methods("GET")
	r.HandleFunc("/carts", createCart).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8002", r))
}
