package db

import (
	"testing"
)

func TestCreate(t *testing.T) {
	d := New()

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
	d := New()

	m := map[string]string{
		"key1": "{name: \"Sean\", age: 50}",
	}
	_ = d.Create(m["key1"])
	t.Run("Create", func(tt *testing.T) {
		index := 2
		if res := d.Update(index, "{name: \"Sean\", age: 35}"); res != true {
			tt.Errorf("not found register with index %d", index)
		}

	})
}
