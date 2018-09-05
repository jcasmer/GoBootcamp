package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type DbInter interface {
	OpenDB() DataBase
	CreateWithIndex(index string, value string) error
	// Create(value string) (string, error)
	Retrieve(index string) (string, error)
	Update(index string, value string) error
	Delete(index string) error
}

type DataBase struct {
	data map[string]string
	mux  sync.Mutex
}

func (db DataBase) CreateWithIndex(index string, value string) error {
	// Método para insertar en la memory db el valor que desea el usuario con la key deseada.
	// Este método valida que el key ingresado este vacio o disponible para la inserción.
	// Si no lo esta devuelve el respectivo mensaje de error

	_, ok := db.data[index]
	if ok {
		return errors.New("No es posible crear el registro con ese indice.")
	}
	db.data[index] = value
	return nil
}

// func (db DataBase) Create(value string) (string, error) {

// 	if db.data[index] != "" {
// 		return errors.New("No es posible crear el registro con el indice. 0. Verifique el archivo o la información a guardar.")
// 	}
// 	if len(db.data) == 0 {
// 		db.data["1"] = value
// 		return "1", nil
// 	}
// 	var keyList []string
// 	for key := range db.data {
// 		keyList = append(keyList, key)
// 	}
// 	sort.Strings(keyList)
// 	// keys := reflect.ValueOf(db.data).MapKeys()
// 	// ktype := keys[len(keys)-1]
// 	// k := ktype.Interface().(int)
// 	k := keyList[len(keyList)-1]
// 	index := string(int32(k) + 1)
// 	db.data[index] = value

// 	fmt.Println("Register created successfully.")
// 	return index, nil
// }

// func (db DataBase) List() bool {

// 	for index, value := range db.data {
// 		fmt.Println(index, value)
// 	}
// 	return true
// }

func (db DataBase) Retrieve(index string) (string, error) {
	// metodo que devuelve el valor de una key dada
	// si el key/indice no existe devuelve el error

	_, ok := db.data[index]
	if ok == false {
		return "", errors.New("NO hay registro con el key indicado")
	}
	return db.data[index], nil
}

func (db DataBase) Update(index string, value string) error {
	// metodo que actualiza el valor de una key dada
	// si el key/indice no existe devuelve el error

	_, ok := db.data[index]
	if ok == false {
		return errors.New("NO hay registro para actualizar con el key indicado")
	}
	db.mux.Lock()
	db.data[index] = value
	db.mux.Unlock()
	fmt.Println("Updated register: ", index, " ", db.data[index])
	return nil
}

func (db DataBase) Delete(index string) error {
	// metodo que actualiza el valor de una key dada
	// si el key/indice no existe devuelve el error

	_, ok := db.data[index]
	if ok == false {
		return errors.New("NO hay registro para eliminar con el key indicado")
	}
	db.mux.Lock()
	delete(db.data, index)
	db.mux.Unlock()
	fmt.Println("Register deleted successfully.")
	return nil
}

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

	b, _ := json.MarshalIndent(db.data, "", " ")

	// writing json to file

	_ = ioutil.WriteFile("db.json", b, 0644)
	return true

}
