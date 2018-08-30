package db

import (
	"fmt"
)

type DbInter interface {
	Create() string
	List() string
	Retrieve() string
	Update() string
	Delete() string
}

type DataBase struct {
	data map[int]string
}

func (db DataBase) Create(value string) string {

	k := len(db.data)
	db.data[k+1] = value
	return "Register created successfully."
}

func (db DataBase) List() string {

	for index, value := range db.data {
		fmt.Println(index, value)
	}
	return ""
}

func (db DataBase) Retrieve(index int) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	fmt.Println("Retrieve:", db.data[index])
	return ""
}

func (db DataBase) Update(index int, value string) string {

	if index < 1 {
		fmt.Println("Not found.")
		return ""
	}
	db.data[index] = value
	fmt.Println("Updated register: ", index, " ", db.data[index])
	return ""
}

func (db DataBase) Delete(index int) string {
	if index < 1 {
		return "Register: Not found."
	}
	delete(db.data, index)
	return "Register deleted successfully."
}

func DbM(db DbInter, method string) {
	if method == "Create" {
		fmt.Println(db.Create())
	}
}

func New() DataBase {
	db := DataBase{make(map[int]string)}
	return db

}

func db() {

	// var db DbInter
	// inserting register
	// p := Person{name: "Sean", age: 50}
	// m := map[string]Person{
	// 	"key1": Person{name: "Sean", age: 50},
	// }
	// out, err := json.Marshal(p)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(p, string(out), m["key1"])
	// dbase := DataBase{"prueba", string(out)}
	// fmt.Println(dbase)
	// DbM(dbase, "Create")
	// //fmt.Println(p.Create("ss", 50))

	// // list all
	// dbase.List()

	// list one
	//p.Retrieve(1)

	// updating a register
	//p.Update(2, "Jhon Doe", 20)

	// delete one
	//fmt.Println(p.Delete(1))

	// show all
	//p.List()

}
