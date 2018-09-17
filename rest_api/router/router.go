package router

import (
	"github.com/gorilla/mux"
	"github.com/jcasmer/GoBootcamp/memory_db/db"
	"github.com/jcasmer/GoBootcamp/rest_api/api/api"
)

// Main function
func Router(db *db.DataBase) *mux.Router {
	// Init router
	r := mux.NewRouter()

	service := api.NewService(db)
	// Route handles & endpoints
	r.HandleFunc("/carts", service.createCart).Methods("POST")
	r.HandleFunc("/carts/{id}", service.getCart).Methods("GET")
	r.HandleFunc("/carts/{id}", service.deleteCart).Methods("DELETE")
	r.HandleFunc("/carts/{id}/items", service.addArticles).Methods("POST")
	r.HandleFunc("/carts/{id}/items", service.deleteAllArticles).Methods("DELETE")
	r.HandleFunc("/carts/{id}/items/{idItem}", service.changeArticles).Methods("PUT")
	r.HandleFunc("/carts/{id}/items/{idItem}", service.deleteArticle).Methods("DELETE")

	return r

}
