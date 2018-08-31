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
	d := db.Open()

	_ = d.Create("prueba")
	fmt.Println(d.List())

	m := map[string]string{
		"key1": "{name: \"juan\", age: 40}",
	}
	// out, err := json.Marshal(p)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(p, string(out), m["key1"])
	_ = d.Create(m["key1"])
	fmt.Println(d.List())
	_ = d.Update(1, "{name: \"Jhon\", age: 50}")
	fmt.Println(d.List())
	d.Close()
}
