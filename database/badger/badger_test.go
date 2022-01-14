package badger

import (
	"fmt"
	"testing"
)

func TestNewSession(t *testing.T) {
	e, err := NewEngine(Option{
		Path: "./data.db",
	})
	if err != nil {
		panic(err)
	}

	err = e.Ts(func(sess ISession) error {
		return sess.SetString("a", "a")
	})
	if err != nil {
		panic(err)
	}

	err = e.Sess().GetString("a", func(v string) error {
		fmt.Println(v)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
