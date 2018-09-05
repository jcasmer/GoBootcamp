package db

import (
	"fmt"
	"testing"
)

var dbName = "db.json"

func TestClose(t *testing.T) {
	db := DataBase{data: make(map[string]string)}

	t.Run("Close", func(tt *testing.T) {
		if res := db.Close(dbName); res != nil {
			tt.Errorf(res.Error())
		}

	})
}

func TestOpen(t *testing.T) {
	// bd := DataBase{data: make(map[string]string)}
	t.Run("Open", func(tt *testing.T) {
		value, res := OpenDB(dbName)
		if res != nil {
			tt.Errorf(res.Error())
		}
		fmt.Println(value)
	})
}

func TestCreate(t *testing.T) {
	d, _ := OpenDB(dbName)
	m := map[string]string{
		"key1": "{name: \"Sean\", age: 40}",
	}
	t.Run("Create", func(tt *testing.T) {
		if res := d.CreateWithIndex("key2", m["key1"]); res != nil {
			tt.Errorf(res.Error())
		}

	})

	d.Close(dbName)
}

func TestRetrieve(t *testing.T) {
	d, _ := OpenDB(dbName)
	t.Run("Create", func(tt *testing.T) {
		value, res := d.Retrieve("key2")
		if res != nil {
			tt.Errorf(res.Error())
		}
		fmt.Println(value)
	})

	d.Close(dbName)
}

func TestUpdate(t *testing.T) {
	d, _ := OpenDB(dbName)
	t.Run("Create", func(tt *testing.T) {
		index := "key4"
		if res := d.Update(index, "{name: \"Sean\", age: 35}"); res != nil {
			tt.Errorf(res.Error())
		}

	})

	d.Close(dbName)
}

func TestDelete(t *testing.T) {
	d, _ := OpenDB(dbName)
	t.Run("Delete", func(tt *testing.T) {
		index := "key2"
		if res := d.Delete(index); res != nil {
			tt.Errorf(res.Error())
		}
	})
	d.Close(dbName)
}
