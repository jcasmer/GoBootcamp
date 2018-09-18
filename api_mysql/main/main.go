package main

import (
	"github.com/jcasmer/GoBootcamp/api_mysql/handler"
)

// Main function
func main() {

	s := handler.Service{}
	s.NewService()
	s.Run(":8002")
}
