package test_test

import (
	"github.com/zhanghup/go-tools/tog/tmp"
	"testing"
)

func TestInsert(t *testing.T) {
	user := User{Id: "11"}

	tmp.Error("涛涛涛涛涛涛涛涛涛涛他")

	t.Run("Insert", func(t *testing.T) {
		err := engine.Sess().Insert(user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := engine.Sess().Table("user").SF("id = ?", "11").Update(map[string]any{
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
