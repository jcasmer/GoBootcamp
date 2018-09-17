package main

import (
	"log"
	"net/http"

	"github.com/jcasmer/GoBootcamp/memory_db/db"
	"github.com/jcasmer/GoBootcamp/rest_api/handler"
)

// Main function
func main() {

	// var err error
	dataBase, _ := db.OpenDB("db.json")
	router := handler.Router(dataBase)
	// service := NewService(dataBase)
	// if err != nil {
	// 	http.Error(, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// Start server
	defer dataBase.Close("db.json")
	log.Fatal(http.ListenAndServe(":8002", router))

	// _ = dataBase.Close("db.json")

}
