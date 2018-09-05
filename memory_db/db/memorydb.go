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
	CreateWithIndex(index string, value string) error
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
	// validamos existencia del key
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
	// validamos existencia del key
	if ok == false {
		return "", errors.New("NO hay registro con el key indicado")
	}
	return db.data[index], nil
}

func (db DataBase) Update(index string, value string) error {
	// metodo que actualiza el valor de una key dada
	// si el key/indice no existe devuelve el error

	_, ok := db.data[index]
	// validamos existencia del key
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
	// metodo que elimina el registro dada una key
	// si el key/indice no existe devuelve el error

	_, ok := db.data[index]
	// validamos existencia del key
	if ok == false {
		return errors.New("NO hay registro para eliminar con el key indicado")
	}
	db.mux.Lock()
	delete(db.data, index)
	db.mux.Unlock()
	fmt.Println("Register deleted successfully.")
	return nil
}

func OpenDB(dbName string) (DataBase, error) {
	// Metodo que permite inicializar la bd.
	// Verifica que haya un archivo si esta realiza la carga de contenido en bd.data
	// Si el archivo no existe lo crea y retorna el objeto bd
	// Si hay errores de lectura o decodificación de la información retorna el error

	db := DataBase{data: make(map[string]string)}

	fileSize, erro := os.Stat(dbName)
	// verificamos existencia.
	if os.IsNotExist(erro) {
		var by []byte
		err := ioutil.WriteFile(dbName, by, 0644)
		if err != nil {
			return DataBase{}, err
		}
	} else {
		// validamos si el archivo no esta vacío para proceder con la lectura y carga de información a db.data
		if fileSize.Size() > 0 {
			// leemos el archivo
			byteValue, err := ioutil.ReadFile(dbName)
			// validamos si hubo un error
			if err != nil {
				return DataBase{}, err
			}
			db.mux.Lock()
			// hacemos un unmarshal para cargar la info del archivo en el puntero de db.data
			err = json.Unmarshal(byteValue, &db.data)
			db.mux.Unlock()
			// validamos si hubo un error
			if err != nil {
				return DataBase{}, err
			}
		}
	}

	return db, nil

}

// func Open(db DbInter) DataBase {
// 	return db.OpenDB()
// }

func (db DataBase) Close(dbName string) error {
	// Método que guarda la información una vez se haya validado una conexión existente con la bd.
	// Retorna errores si se encuentran con los respectivos mensajes

	//se valida si hay la data contiene información
	if len(db.data) == 0 {
		return errors.New("No hay conexión con la bd. Por favor conectese a la bd y luego proceda a cerrarla.")
	}

	b, erro := json.MarshalIndent(db.data, "", " ")
	//se valida si hay algún error en el encode to json
	if erro != nil {
		return errors.New("Verifique el formato de la información. Debe ser (map)")
	}
	// escribimos el archivo.
	err := ioutil.WriteFile(dbName, b, 0644)
	//se valida si hay algún error en las escritura del archivo
	if err != nil {
		return errors.New("Error.Verifique la información del archivo ó si no cuenta con los permisos necesarios para escribir")
	}
	return nil
}
