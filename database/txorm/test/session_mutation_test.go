package test_test

import (
	"github.com/zhanghup/go-tools/tog"
	"testing"
)

func TestInsert(t *testing.T) {
	user := User{
		Id:   "111",
		Name: "111",
	}

	tog.Error("涛涛涛涛涛涛涛涛涛涛他")

	t.Run("Insert", func(t *testing.T) {
		err := engine.Session(true).Insert(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := engine.Session(true).Table("user").SF("id = ?", "11").Update(map[string]interface{}{
			"name": "112",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := engine.Session(true).Table("user").SF("id = ?", "11").Delete()
		if err != nil {
			t.Fatal(err)
		}
	})

}
