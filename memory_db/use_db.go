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

	m := map[string]string{
		"key1": "{name: \"Sean\", age: 50}",
	}
	// out, err := json.Marshal(p)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(p, string(out), m["key1"])
	fmt.Println(d.Create(m["key1"]))
	fmt.Println(d.List())
	d.Update(1, "{name: \"Jhon\", age: 50}")
	fmt.Println(d.List())

}
