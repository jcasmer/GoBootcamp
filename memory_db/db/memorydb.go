package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
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
	mux  sync.Mutex
}

func (db DataBase) Create(value string) DataBase {
	db.mux.Lock()
	k := len(db.data)
	db.data[k+1] = value
	fmt.Println("Register created successfully.")
	db.mux.Unlock()
	return db
}

func (db DataBase) List() string {

	db.mux.Lock()
	defer db.mux.Unlock()
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
	db.mux.Lock()
	defer db.mux.Unlock()
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
	db.mux.Lock()
	db.data[index] = value
	fmt.Println("Updated register: ", index, " ", db.data[index])
	db.mux.Unlock()
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
	db.mux.Lock()
	delete(db.data, index)
	fmt.Println("Register deleted successfully.")
	db.mux.Unlock()
	return true
}

// func DbM(db DbInter, method string) {
// 	if method == "Create" {
// 		fmt.Println(db.Create())
// 	}
// }

func Open() DataBase {
	db := DataBase{data: make(map[int]string)}

	byteValue, _ := ioutil.ReadFile("db.json")
	err := json.Unmarshal(byteValue, &db.data)
	if err != nil {
		db = DataBase{data: make(map[int]string)}
	}
	else 
	return db

}

func (db DataBase) Close() bool {

	// printing out json neatly to demonstrate
	b, _ := json.MarshalIndent(db.data, "", " ")

	// writing json to file

	_ = ioutil.WriteFile("db.json", b, 0644)
	return true

}
