package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

var dataPerson = make(map[int]Person)

func (Person) Create(name string, age int) string {

	k := len(dataPerson)
	dataPerson[k+1] = Person{name, age}
	return "Person was created successfully."
}

func (Person) List() {

	for _, value := range dataPerson {
		fmt.Println(value)
	}
}

func (Person) Retrieve(index int) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	fmt.Println("Retrieve:", dataPerson[index])
	return ""
}

func (Person) Update(index int, name string, age int) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	dataPerson[index] = Person{name, age}
	fmt.Println("Updated register: ", index, " ", dataPerson[index])
	return ""
}

func (Person) Delete(index int) string {
	if index < 1 {
		return "Register: Not found."
	}
	delete(dataPerson, index)
	return "Person was deleted successfully."
}

func main() {

	// inserting register
	p := Person{}
	fmt.Println(p.Create("Sean", 50))
	fmt.Println(p.Create("ss", 50))

	// list all
	p.List()

	// list one
	p.Retrieve(1)

	// updating a register
	p.Update(2, "Jhon Doe", 20)

	// delete one
	fmt.Println(p.Delete(1))

	// show all
	p.List()

}
