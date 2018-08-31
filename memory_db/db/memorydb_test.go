package db

import (
	"testing"
)

func TestCreate(t *testing.T) {
	d := Open()

	m := map[string]string{
		"key1": "{name: \"Sean\", age: 50}",
	}
	t.Run("Create", func(tt *testing.T) {
		if res := d.Create(m["key1"]); res != true {
			tt.Errorf("expected %s but got ", string(m["key1"]))
		}

	})
}

func TestUpdate(t *testing.T) {
	d := Open()

	m := map[string]string{
		"key1": "{name: \"Sean\", age: 50}",
	}
	_ = d.Create(m["key1"])
	t.Run("Create", func(tt *testing.T) {
		index := 1
		if res := d.Update(index, "{name: \"Sean\", age: 35}"); res != true {
			tt.Errorf("not found register with index %d to update", index)
		}

	})
}

func TestDelete(t *testing.T) {
	d := Open()

	m := map[string]string{
		"key1": "{name: \"Sean\", age: 50}",
	}
	_ = d.Create(m["key1"])
	t.Run("Delete", func(tt *testing.T) {
		index := 1
		if res := d.Delete(index); res != true {
			tt.Errorf("not found register with index %d to Delete", index)
		}

	})
	// t.Run("Delete", func(tt *testing.T) {
	// 	index := 1
	// 	if res := d.Delete(index); res != true {
	// 		tt.Errorf("not found register with index %d to Delete", index)
	// 	}

	// })
}
