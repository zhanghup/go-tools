package test_test

import (
	"testing"
)

func TestInsert(t *testing.T) {
	user := User{"11", "11", 11}

	t.Run("Insert", func(t *testing.T) {
		err := engine.SessionAuto().Insert(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := engine.SessionAuto().Table("user").SF("where id = ?", "11").Update(map[string]interface{}{
			"name": "112",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := engine.SessionAuto().Table("user").SF("where id = ?", "11").Delete()
		if err != nil {
			t.Fatal(err)
		}
	})

}
