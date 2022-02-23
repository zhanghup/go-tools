package test_test

import (
	"github.com/zhanghup/go-tools/tog"
	"testing"
)

func TestInsert(t *testing.T) {
<<<<<<< HEAD
	user := User{"11", "11", 11, nil, nil, nil}
=======
	user := User{
		Id:   "111",
		Name: "111",
	}
>>>>>>> c64ebceac3813b7f2621365310fb7b47ecc4b4b8

	tog.Error("涛涛涛涛涛涛涛涛涛涛他")

	t.Run("Insert", func(t *testing.T) {
		err := engine.Sess().Insert(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := engine.Sess().Table("user").SF("id = ?", "11").Update(map[string]interface{}{
			"name": "112",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := engine.Sess().Table("user").SF("id = ?", "11").Delete()
		if err != nil {
			t.Fatal(err)
		}
	})

}
