package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"
)

type DbInter interface {
	// Open() DataBase
	// close()
	OpenDB() DataBase
	Create(value string) bool
	List() string
	Retrieve(index int) string
	Update(index int, value string) bool
	Delete(index int) bool
}

type DataBase struct {
	data map[int]string
	mux  sync.Mutex
}

func (db DataBase) Create(value string) bool {

	db.mux.Lock()
	if len(db.data) == 0 {
		db.data[1] = value
	} else {
		var keyList []int
		for key := range db.data {
			keyList = append(keyList, key)
		}
		sort.Ints(keyList)
		// keys := reflect.ValueOf(db.data).MapKeys()
		// ktype := keys[len(keys)-1]
		// k := ktype.Interface().(int)
		k := keyList[len(keyList)-1]
		db.data[k+1] = value
	}
	fmt.Println("Register created successfully.")
	db.mux.Unlock()
	return true
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

func (db DataBase) OpenDB() DataBase {

	// db = DataBase{data: make(map[int]string)}
	db.mux.Lock()
	// db.data = make(map[int]string)
	fileSize, erro := os.Stat("db.json")
	if os.IsNotExist(erro) {
		var by []byte
		_ = ioutil.WriteFile("db.json", by, 0644)
	} else {
		if fileSize.Size() > 0 {
			byteValue, _ := ioutil.ReadFile("db.json")

			err := json.Unmarshal(byteValue, &db.data)
			if err != nil {
				// db.data = make(map[int]string)
				return DataBase{}
				// db := DataBase{data: make(map[int]string)}
				// return false
			}
		}
	}
	db.mux.Unlock()
	return db

}

func Open(db DbInter) DataBase {
	return db.OpenDB()
}

func (db DataBase) Close() bool {

	// printing out json neatly to demonstrate
	b, _ := json.MarshalIndent(db.data, "", " ")

	// writing json to file

	_ = ioutil.WriteFile("db.json", b, 0644)
	return true

}
