package test_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestTables(t *testing.T) {
	fmt.Println(tools.JSONString(engine.Tables(), true))
}

func TestTable(t *testing.T) {
	fmt.Println(tools.JSONString(engine.Table("user"), true))
}

func TestColumn(t *testing.T) {
	fmt.Println(tools.JSONString(engine.Table("user").Column("id"), true))
}
