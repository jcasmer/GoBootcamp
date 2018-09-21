package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jcasmer/GoBootcamp/memory_db/dbMysql"
	"github.com/jcasmer/GoBootcamp/rest_api/handler"
)

// Main function
func main() {

	db, _ :=  dbMysql.OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	s := handler.NewService(db)
	r := mux.NewRouter()
	r.PathPrefix("/api").Handler(s)
	log.Fatal(http.ListenAndServe(":8002", r))
}
