package handler

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
	r.HandleFunc("/carts", service.CreateCart).Methods("POST")
	r.HandleFunc("/carts/{id}", service.GetCart).Methods("GET")
	r.HandleFunc("/carts/{id}", service.deleteCart).Methods("DELETE")
	r.HandleFunc("/carts/{id}/items", service.AddArticles).Methods("POST")
	r.HandleFunc("/carts/{id}/items", service.deleteAllArticles).Methods("DELETE")
	r.HandleFunc("/carts/{id}/items/{idItem}", service.changeArticles).Methods("PUT")
	r.HandleFunc("/carts/{id}/items/{idItem}", service.deleteArticle).Methods("DELETE")

	return r

}
