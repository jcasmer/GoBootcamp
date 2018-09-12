package main

import (
	"log"
	"net/http"

	"github.com/jcasmer/GoBootcamp/memory_db/db"
)

// Main function
func main() {

	// var err error
	dataBase, _ := db.OpenDB("db.json")
	router := Router(dataBase)
	// if err != nil {
	// 	http.Error(, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// Start server
	// defer dataBase.Close("db.json")
	log.Fatal(http.ListenAndServe(":8002", router))
	_ = dataBase.Close("db.json")

}
