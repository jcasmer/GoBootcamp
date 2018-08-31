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

func (db DataBase) Create(value string) bool {

	k := len(db.data)
	db.data[k+1] = value
	fmt.Println("Register created successfully.")
	return true
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

func (db DataBase) Update(index int, value string) bool {

	if index < 1 {
		fmt.Println("Not found.")
		return false
	}
	if db.data[index] == "" {
		fmt.Println("Not found.")
		return false
	}
	db.data[index] = value
	fmt.Println("Updated register: ", index, " ", db.data[index])
	return true
}

func (db DataBase) Delete(index int) bool {
	if index < 1 {
		fmt.Println("Register: Not found.")
		return false
	}
	if db.data[index] == "" {
		fmt.Println("Register: Not found.")
		return false
	}
	delete(db.data, index)
	fmt.Println("Register deleted successfully.")
	return true
}

// func DbM(db DbInter, method string) {
// 	if method == "Create" {
// 		fmt.Println(db.Create())
// 	}
// }

func New() DataBase {
	db := DataBase{make(map[int]string)}
	return db

}
