package loader_test

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools/loader"
	"testing"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id   string `json:"id" xorm:"pk"`
	Type string `json:"type"`
}

func TestLoaderObject(t *testing.T) {
	loader := loader.NewObjectLoader[int](func(keys []string) (map[string]int, error) {
		fmt.Println("111111111111111111111111")
		return map[string]int{
			"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
			"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		}, nil
	})

	for i := 0; i < 10; i++ {
		go func(n int) {
			v, _, err := loader.Load(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoader(t *testing.T) {
	loader := loader.Load("123", func(keys []string) (map[string]int, error) {
		fmt.Println("111111111111111111111111")
		return map[string]int{
			"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
			"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		}, nil
	})
	for i := 0; i < 10; i++ {
		go func(n int) {
			v, _, err := loader.Load(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoadSlice(t *testing.T) {
	loader := loader.Slice[User](nil, "user", "type")

	for i := 0; i < 3; i++ {
		go func(n int) {
			v, err := loader(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoadSlice2(t *testing.T) {
	loader := loader.Slice[User](nil, "select * from user where type in :keys", "type")

	for i := 0; i < 3; i++ {
		go func(n int) {
			v, err := loader(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoadObject(t *testing.T) {
	loader := loader.Info[User](nil, "user", "id")

	for i := 0; i < 600; i++ {
		go func(n int) {
			v, err := loader(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoadObject2(t *testing.T) {
	loader := loader.Info[User](nil, "select * from user where user.id in :keys", "id")

	for i := 0; i < 600; i++ {
		go func(n int) {
			v, err := loader(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestLoadObject3(t *testing.T) {
	loader := loader.InfoId[User](nil, "user")

	for i := 0; i < 10; i++ {
		go func(n int) {
			v, err := loader(fmt.Sprintf("%d", n))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("vvvvvvvvv: ", v)
		}(i)
	}
	time.Sleep(time.Second)
}

func init() {
	db, err := xorm.NewEngine("sqlite3", "./data.dat")
	if err != nil {
		panic(err)
	}
	db.ShowSQL(true)
	loader.SetDB(db)
	db.DropTables(User{})
	db.Sync2(User{})

	for i := 0; i < 10; i++ {
		db.InsertOne(User{
			Id:   fmt.Sprintf("%d", i),
			Type: fmt.Sprintf("%d", i%3),
		})
	}

}
