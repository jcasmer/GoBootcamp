package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	// d := db.Open()

	// _ = d.Create("prueba")
	// fmt.Println(d.List())

	m := map[int]Person{
		1: Person{"daniel", 15},
		2: Person{"juan", 30},
	}
	v := make(map[int]string)
	fmt.Println(m[1])
	v[1] = {string(m[1].name), string(m[1].age)}
	fmt.Println(v[1])
	// _ = d.Create(string(v[1]))
	// // fmt.Println(d.List())
	// // _ = d.Update(1, "{name: \"Jhon\", age: 50}")
	// fmt.Println(d.List())
	// d.Close()
}
