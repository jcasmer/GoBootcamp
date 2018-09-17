package main

import (
	"github.com/jcasmer/GoBootcamp/rest_api/handler"
)

// Main function
func main() {

	s := handler.Service{}
	s.NewService()
	s.Run(":8002")
}
