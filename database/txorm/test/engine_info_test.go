package test_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestTables(t *testing.T) {
	infos, err := engine.Tables()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tools.JSONString(infos, true))
}

func TestTable(t *testing.T) {
	info, err := engine.Table("user")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tools.JSONString(info, true))
}

func TestColumn(t *testing.T) {
	info, err := engine.Table("user")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tools.JSONString(info.Column("id"), true))
}