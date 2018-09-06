package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var (

	//ErrNotFound es retornado cuando no existe el indice o el item en la bd
	ErrNotFound = errors.New("not found")

	// ErrInvalidDataBase es retornado cuando el archivo de la bd es invalida
	ErrInvalidDataBase = errors.New("invalid database")

	// ErrInvalidFormat es retornado cuando hay un formato inválido
	ErrInvalidFormat = errors.New("invalid format")

	// ErrPerms es retonado cuando no cuenta con los permisos necesarios
	ErrPerms = errors.New("unauthorized")

	//ErrDatabaseClosed es retornado cuando la base de datos esta cerrada y se esta tratando de hacer operaciones
	ErrDatabaseClosed = errors.New("database closed")

	// ErrIndexExists es retornado cuando el indice del objeto a crear ya existe
	ErrIndexExists = errors.New("index exists")
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
	open bool
}

func (db *DataBase) CreateWithIndex(index string, value string) error {
	// Método para insertar en la memory db el valor que desea el usuario con la key deseada.
	// Este método valida que el key ingresado este vacio o disponible para la inserción.
	// Si no lo esta devuelve el respectivo mensaje de error

	// validamos que haya conexión con la bd
	if db.open == false {
		return ErrDatabaseClosed
	}
	_, ok := db.data[index]
	// validamos existencia del key
	if ok {
		return ErrIndexExists
	}
	db.data[index] = value
	return nil
}

func (db *DataBase) Retrieve(index string) (string, error) {
	// metodo que devuelve el valor de una key dada
	// si el key/indice no existe devuelve el error

	// validamos que haya conexión con la bd
	if db.open == false {
		return "", ErrDatabaseClosed
	}

	_, ok := db.data[index]
	// validamos existencia del key
	if ok == false {
		return "", ErrNotFound
	}
	return db.data[index], nil
}

func (db *DataBase) Update(index string, value string) error {
	// metodo que actualiza el valor de una key dada
	// si el key/indice no existe devuelve el error

	// validamos que haya conexión con la bd
	if db.open == false {
		return ErrDatabaseClosed
	}

	_, ok := db.data[index]
	// validamos existencia del key
	if ok == false {
		return ErrNotFound
	}
	db.mux.Lock()
	db.data[index] = value
	db.mux.Unlock()
	fmt.Println("Updated register: ", index, " ", db.data[index])
	return nil
}

func (db *DataBase) Delete(index string) error {
	// metodo que elimina el registro dada una key
	// si el key/indice no existe devuelve el error

	// validamos que haya conexión con la bd
	if db.open == false {
		return ErrDatabaseClosed
	}

	_, ok := db.data[index]
	// validamos existencia del key
	if ok == false {
		return ErrNotFound
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
				return DataBase{}, ErrInvalidDataBase
			}
			db.mux.Lock()
			// hacemos un unmarshal para cargar la info del archivo en el puntero de db.data
			err = json.Unmarshal(byteValue, &db.data)
			db.mux.Unlock()
			// validamos si hubo un error
			if err != nil {
				return DataBase{}, ErrInvalidFormat
			}
			db.open = true
		}
	}

	return db, nil

}

// func Open(db DbInter) DataBase {
// 	return db.OpenDB()
// }

func (db *DataBase) Close(dbName string) error {
	// Método que guarda la información una vez se haya validado una conexión existente con la bd.
	// Retorna errores si se encuentran con los respectivos mensajes
	db.mux.Lock()
	defer db.mux.Unlock()
	if db.open == false {
		return ErrDatabaseClosed
	}

	b, erro := json.MarshalIndent(db.data, "", " ")
	//se valida si hay algún error en el encode to json
	if erro != nil {
		return ErrInvalidFormat
	}
	// escribimos el archivo.
	err := ioutil.WriteFile(dbName, b, 0644)
	//se valida si hay algún error en las escritura del archivo
	if err != nil {
		return ErrPerms
	}
	db.open = false
	return nil
}
