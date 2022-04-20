package test_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestGet(t *testing.T) {
	user := User{}
	_, err := engine.Sess().SF(" id = ?", "1").Get(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(user))
}

func TestFind(t *testing.T) {
	user := make([]User, 0)
	err := engine.Sess().SF("limit 10").Find(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(user))
}

func TestCount(t *testing.T) {
	t.Run("where", func(t *testing.T) {
		n, err := engine.Sess().Table("user").SF(" id = ?", "1").Count()
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Fatal("错误")
		}
	})

	t.Run("all", func(t *testing.T) {
		n, err := engine.Sess().Table("user").Count()
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})
}

func TestExt(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		n, err := engine.Sess().SF("select count(1) from user").Int()
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})

	t.Run("Int64", func(t *testing.T) {
		n, err := engine.Sess().SF("select count(1) from user").Int64()
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})

	t.Run("Float64", func(t *testing.T) {
		n, err := engine.Sess().SF("select count(1) from user").Float64()
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})

	t.Run("String", func(t *testing.T) {
		n, err := engine.Sess().SF("select count(1) from user").String()
		if err != nil {
			t.Fatal(err)
		}
		if n != "10" {
			t.Fatal("错误")
		}
	})

	t.Run("Strings", func(t *testing.T) {
		n, err := engine.Sess().SF("select id from user limit 4").Strings()
		if err != nil {
			t.Fatal(err)
		}
		if len(n) != 4 {
			t.Fatal("错误")
		}
	})

}

func TestPage(t *testing.T) {
	t.Run("Page Size < 0 ", func(t *testing.T) {
		users := make([]User, 0)
		_, err := engine.Sess().Page(1, -1, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size < 0 with select", func(t *testing.T) {
		users := make([]User, 0)
		_, err := engine.Sess().SF("select * from user where 1 = 1").Page(1, -1, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size < 0 with select2", func(t *testing.T) {
		users := make([]User, 0)
		_, err := engine.Sess().SF(" 1 = 1").Page(1, -1, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size < 0 with select3", func(t *testing.T) {
		users := make([]User, 0)
		_, err := engine.Sess().SF(" id = ?", "1").Page(1, -1, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 1 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size < 0 with select4", func(t *testing.T) {
		users := make([]User, 0)
		_, err := engine.Sess().SF(" id = :id", map[string]any{"id": "1"}).Page(1, -1, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 1 {
			t.Fatal("错误")
		}
	})

	t.Run("Page Size = 0 ", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().Page(1, 0, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size = 0 with select", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF("select * from user where 1 = 1").Page(1, 0, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size = 0 with select2", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" 1 = 1").Page(1, 0, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size = 0 with select3", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" id = ?", "1").Page(1, 0, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size = 0 with select4", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" id = :id", map[string]any{"id": "1"}).Page(1, 0, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Fatal("错误")
		}
	})

	t.Run("Page Size > 0", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tools.JSONString(users))
		fmt.Println(n)
	})

	t.Run("Page Size > 0 ", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 || len(users) != 4 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size > 0 with select", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF("select * from user where 1 = 1").Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 || len(users) != 4 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size > 0 with select2", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" 1 = 1").Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 10 || len(users) != 4 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size > 0 with select3", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" id = ?", "1").Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 || len(users) != 1 {
			t.Fatal("错误")
		}
	})
	t.Run("Page Size > 0 with select4", func(t *testing.T) {
		users := make([]User, 0)
		n, err := engine.Sess().SF(" id = :id", map[string]any{"id": "1"}).Page(1, 4, true, &users)
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 || len(users) != 1 {
			t.Fatal("错误")
		}
	})

}
