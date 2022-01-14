package buntdb

import "github.com/tidwall/buntdb"

func (this *Engine) CreateIndex(name, pattern string, less ...func(a, b string) bool) error {
	return this.db.CreateIndex(name, pattern, less...)
}
func (this *Engine) CreateIndexSignal(name, pattern string, less func(a, b string) bool, desc ...bool) error {
	idx := less
	if len(desc) > 0 && desc[0] {
		idx = buntdb.Desc(idx)
	}
	return this.CreateIndex(name, pattern, idx)
}

func (this *Engine) CreateSpatialIndex(name, pattern string, rect func(item string) (min, max []float64)) error {
	return this.db.CreateSpatialIndex(name, pattern, rect)
}

func (this *Engine) CreateStringIndex(name, pattern string, desc ...bool) error {
	return this.CreateIndexSignal(name, pattern, buntdb.IndexString, desc...)
}
func (this *Engine) CreateBinaryIndex(name, pattern string, desc ...bool) error {
	return this.CreateIndexSignal(name, pattern, buntdb.IndexBinary, desc...)
}
func (this *Engine) CreateIntIndex(name, pattern string, desc ...bool) error {
	return this.CreateIndexSignal(name, pattern, buntdb.IndexInt, desc...)
}
func (this *Engine) CreateUintIndex(name, pattern string, desc ...bool) error {
	return this.CreateIndexSignal(name, pattern, buntdb.IndexUint, desc...)
}
func (this *Engine) CreateFloatIndex(name, pattern string, desc ...bool) error {
	return this.CreateIndexSignal(name, pattern, buntdb.IndexFloat, desc...)
}
