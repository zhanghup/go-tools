package buntdb

import (
	"github.com/tidwall/buntdb"
	"strings"
)

func (this *Engine) IndexSpatialCreate(name, pattern string, rect func(item string) (min, max []float64)) error {
	return this.db.CreateSpatialIndex(name, pattern, rect)
}

/*
	IndexCreate 通用索引创建

	index类型包含: string/binary/int/uint/float
	例如：
		indexes = ["string","string:asc","string:desc"]
*/
func (this *Engine) IndexCreate(name, pattern string, indexes ...string) error {
	idxs := make([]func(a, b string) bool, 0)

	for _, s := range indexes {
		ss := strings.Split(s, ":")
		var v func(a, b string) bool
		switch ss[0] {
		case "string":
			v = buntdb.IndexString
		case "binary":
			v = buntdb.IndexBinary
		case "int":
			v = buntdb.IndexInt
		case "uint":
			v = buntdb.IndexUint
		case "float":
			v = buntdb.IndexFloat
		}

		if len(ss) == 2 && ss[1] == "desc" {
			v = buntdb.Desc(v)
		}
		idxs = append(idxs, v)
	}

	return this.db.CreateIndex(name, pattern, idxs...)
}

/*
	IndexJsonCreate JSON索引创建
	indexes = ["user.age","user.age:asc","user.age:desc]
*/
func (this *Engine) IndexJsonCreate(name, pattern string, indexes ...string) error {
	idxs := make([]func(a, b string) bool, 0)

	for _, s := range indexes {
		ss := strings.Split(s, ":")
		if len(ss) == 2 && ss[1] == "desc" {
			idxs = append(idxs, buntdb.Desc(buntdb.IndexJSON(ss[0])))
		} else {
			idxs = append(idxs, buntdb.IndexJSON(ss[0]))
		}
	}

	return this.db.CreateIndex(name, pattern, idxs...)
}

func (this *Engine) IndexRectCreate(name, pattern string) error {
	return this.IndexSpatialCreate(name, pattern, buntdb.IndexRect)
}

func (this *Engine) Indexes() ([]string, error) {
	return this.db.Indexes()
}
func (this *Engine) IndexDrop(name string) error {
	return this.db.DropIndex(name)
}
