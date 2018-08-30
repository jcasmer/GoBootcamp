package main

import (
	"fmt"

	db "./db"
)

type Person struct {
	name string
	age  int
}

func main() {
	d := db.New()

	p := Person{name: "Sean", age: 50}
	fmt.Println(d.Create("prueba"), p)
	fmt.Println(d.List())

}
