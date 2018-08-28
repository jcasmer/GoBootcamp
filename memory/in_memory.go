package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

var dataPerson = make(map[int]Person)

func Create(name string, age int) string {

	k := len(dataPerson)
	dataPerson[k+1] = Person{name, age}
	return "Person was created successfully."
}

func Retrieve(index int) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	fmt.Println("Retrieve:", dataPerson[index])
	return ""
}

func Update(index int, name string, age int) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	dataPerson[index] = Person{name, age}
	fmt.Println("Updated register: ", index, " ", dataPerson[index])
	return ""
}

func Delete(index int) string {
	if index < 1 {
		return "Register: Not found."
	}
	delete(dataPerson, index)
	return "Person was deleted successfully."
}

func main() {

	// inserting register
	fmt.Println(Create("Sean", 50))
	fmt.Println(Create("ss", 50))

	// for _, value := range dataPerson {
	// 	fmt.Println(value)
	// }

	// list one
	Retrieve(1)

	// updating a register
	Update(2, "Jhon Doe", 20)

	// delete one
	fmt.Println(Delete(1))

	// show all
	for _, value := range dataPerson {
		fmt.Println(value)
	}

}
