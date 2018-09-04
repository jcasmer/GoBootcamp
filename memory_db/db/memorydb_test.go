package db

import (
	"testing"
)

// func TestOpen(t *testing.T) {
// 	t.Run("Open", func(tt *testing.T) {
// 		d := Open()
// 		if len(d.data) == 0 {
// 			tt.Errorf("Failed to open db. Please check the file")
// 		}

// 	})
// }

func TestCreate(t *testing.T) {
	bd := DataBase{data: make(map[int]string)}
	d := Open(bd)
	// fmt.Println(d)
	m := map[string]string{
		"key1": "{name: \"Sean\", age: 40}",
	}
	t.Run("Create", func(tt *testing.T) {
		if res := d.Create(m["key1"]); res != true {
			tt.Errorf("expected %s but got ", string(m["key1"]))
		}

	})

	d.Close()
}

func TestUpdate(t *testing.T) {
	bd := DataBase{data: make(map[int]string)}
	d := Open(bd)
	// m := map[string]string{
	// 	"key1": "{name: \"Sean\", age: 30}",
	// }
	// _ = d.Create(m["key1"])
	t.Run("Create", func(tt *testing.T) {
		index := 2
		if res := d.Update(index, "{name: \"Sean\", age: 35}"); res != true {
			tt.Errorf("not found register with index %d to update", index)
		}

	})

	d.Close()
}

// func TestDelete(t *testing.T) {
// 	d := Open()
// 	m := map[string]string{
// 		"key1": "{name: \"Sean\", age: 50}",
// 	}
// 	_ = d.Create(m["key1"])
// 	t.Run("Delete", func(tt *testing.T) {
// 		index := 1
// 		if res := d.Delete(index); res != true {
// 			tt.Errorf("not found register with index %d to Delete", index)
// 		}

// 	})
// 	d.Close()
// }

// func TestClose(t *testing.T) {

// 	t.Run("Delete", func(tt *testing.T) {
// 		if res :=d.Close()(); res != true {
// 			tt.Errorf("Close error ")
// 		}

// 	})
// }
