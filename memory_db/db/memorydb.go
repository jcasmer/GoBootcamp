package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (db DataBase) Create(value string) DataBase {

	k := len(db.data)
	db.data[k+1] = value
	fmt.Println("Register created successfully.")
	return db
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

func Open() DataBase {
	db := DataBase{make(map[int]string)}

	byteValue, _ := ioutil.ReadFile("db.json")
	err := json.Unmarshal(byteValue, &db.data)
	if err != nil {
		// nozzle.printError("opening config file", err.Error())
	}
	return db

}

func (db DataBase) Close() bool {

	// printing out json neatly to demonstrate
	b, _ := json.MarshalIndent(db.data, "", " ")

	// writing json to file

	_ = ioutil.WriteFile("db.json", b, 0644)
	return true
	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	// f, err := os.OpenFile("/var/log/debug-web.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // write to file, f.Write()
	// f.Write(b)

}
