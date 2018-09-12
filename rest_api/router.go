package main

import (
	"github.com/gorilla/mux"
	"github.com/jcasmer/GoBootcamp/memory_db/db"
)

// Main function
func Router(db *db.DataBase) *mux.Router {
	// Init router
	r := mux.NewRouter()

	service := NewService(db)
	// Route handles & endpoints
	r.HandleFunc("/carts", service.createCart).Methods("POST")
	r.HandleFunc("/carts/{id}", service.getCart).Methods("GET")
	// r.HandleFunc("/carts/{id}/items", service.addArticles).Methods("POST")

	return r

}
