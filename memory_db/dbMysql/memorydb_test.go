package dbMysql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var dbName = "db.json"

// func TestClose(t *testing.T) {
// 	// test cuando no se ha abierto una conexion a la bd
// 	db := DataBase{data: make(map[string]string)}

// 	t.Run("Close", func(tt *testing.T) {
// 		if res := db.Close(dbName); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})
// }

// func TestCloseSuccess(t *testing.T) {
// 	// test cuando no si ha abierto una conexion a la bd
// 	d, _ := OpenDB(dbName)

// 	t.Run("Close", func(tt *testing.T) {
// 		if res := d.Close(dbName); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})
// }

func TestOpen(t *testing.T) {
	// test exito de una conexión abierta correctamente
	t.Run("Open", func(tt *testing.T) {
		// db, err := sql.Open("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
		d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		var id int
		var data string
		statement := fmt.Sprintf("SELECT id, data FROM tests WHERE id=1")
		err = d.db.QueryRow(statement).Scan(&id, &data)
		fmt.Println(id, data)

	})
}

// func TestOpenFail(t *testing.T) {
// 	// test fall de una conexión abierta correctamente
// 	t.Run("Open Fail", func(tt *testing.T) {
// 		_, res := OpenDB("someName")
// 		if res != nil {
// 			tt.Errorf(res.Error())
// 		}
// 	})
// }

func TestCreate(t *testing.T) {
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Create", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		if _, res := d.CreateWithIndex("{name: \"Sean\", age: 40}"); res != nil {
			tt.Errorf(res.Error())
		}

	})
	d.Close(dbName)
}

// func TestCreateCloseDB(t *testing.T) {
// 	d := DataBase{data: make(map[string]string)}
// 	m := map[string]string{
// 		"key1": "{name: \"Sean\", age: 40}",
// 	}
// 	t.Run("Create", func(tt *testing.T) {
// 		if res := d.CreateWithIndex("key2", m["key1"]); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})

// 	d.Close(dbName)
// }

// func TestCreateFailIndex(t *testing.T) {
// 	d := DataBase{data: make(map[string]string)}
// 	m := map[string]string{
// 		"key1": "{name: \"Sean\", age: 40}",
// 	}
// 	t.Run("Create", func(tt *testing.T) {
// 		if res := d.CreateWithIndex("key1", m["key1"]); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})

// 	d.Close(dbName)
// }

// func TestCreateFailDb(t *testing.T) {
// 	// test de actualizar un registro cuando no hay conexión con la bd
// 	d := DataBase{data: make(map[string]string)}
// 	m := map[string]string{
// 		"key1": "{name: \"Sean\", age: 40}",
// 	}
// 	t.Run("Create", func(tt *testing.T) {
// 		if res := d.CreateWithIndex("key1", m["key1"]); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})

// 	d.Close(dbName)
// }

func TestRetrieveFail(t *testing.T) {
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Retrieve", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		value, res := d.Retrieve("2")
		if res != nil {
			tt.Errorf(res.Error())
		}
		fmt.Println(value)
	})

	d.Close(dbName)
}

func TestRetrieve(t *testing.T) {
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Retrieve", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		value, res := d.Retrieve("11")
		if res != nil {
			tt.Errorf(res.Error())
		}
		fmt.Println(value)
	})

	d.Close(dbName)
}

func TestUpdate(t *testing.T) {
	// test para actualizar un registro
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Update", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		index := "8"
		if res := d.Update(index, "{name: \"Sean\", age: 34}"); res != nil {
			tt.Errorf(res.Error())
		}

	})
	d.Close(dbName)
}

func TestUpdateFail(t *testing.T) {
	// test para actualizar un registro con un indice inexistente
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Update", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		index := "2"
		if res := d.Update(index, "{name: \"Sean\", age: 34}"); res != nil {
			tt.Errorf(res.Error())
		}

	})
	d.Close(dbName)
}

func TestDeleteFail(t *testing.T) {
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Delete Fail", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		index := "2"
		if res := d.Delete(index); res != nil {
			tt.Errorf(res.Error())
		}
	})
	d.Close(dbName)
}

func TestDelete(t *testing.T) {
	d, err := OpenDB("mysql", "root:k4tt14n4**@tcp(172.17.0.4:3306)/GoBootcamp")
	t.Run("Delete Fail", func(tt *testing.T) {
		if err != nil {
			tt.Errorf(err.Error())
			return
		}
		index := "5"
		if res := d.Delete(index); res != nil {
			tt.Errorf(res.Error())
		}
	})
	d.Close(dbName)
}

// func TestDeleteFailDb(t *testing.T) {
// 	// test de eliminar un registro cuando no hay conexión con la bd
// 	d := DataBase{data: make(map[string]string)}
// 	t.Run("Delete Fail Close DB", func(tt *testing.T) {
// 		index := "key4"
// 		if res := d.Delete(index); res != nil {
// 			tt.Errorf(res.Error())
// 		}

// 	})

// 	d.Close(dbName)
// }
