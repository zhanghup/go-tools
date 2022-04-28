package influx

import (
	"fmt"
	"strings"
)

type QueryString struct {
	engine *Engine
	data   []string
}

func (this QueryString) Query(s string) QueryString {
	this.data = append(this.data, s)
	return this
}
func (this QueryString) Range1(start string) QueryString {
	this.data = append(this.data, fmt.Sprintf(`range(start: %s)`, start))
	return this
}
func (this QueryString) Range2(stop string) QueryString {
	this.data = append(this.data, fmt.Sprintf(`range(stop: %s)`, stop))
	return this
}
func (this QueryString) Range3(start, stop string) QueryString {
	this.data = append(this.data, fmt.Sprintf(`range(start: %s, stop: %s)`, start, stop))
	return this
}

func (this QueryString) Filter(fn string, dropEntityData ...bool) QueryString {
	onEmpty := "drop"
	if len(dropEntityData) > 0 && !dropEntityData[0] {
		onEmpty = "keep"
	}
	this.data = append(this.data, fmt.Sprintf(`filter(fn: %s, onEmpty: %s)`, fn, onEmpty))
	return this
}

func (this QueryString) Limit1(limit int) QueryString {
	this.data = append(this.data, fmt.Sprintf(`limit(n: %d)`, limit))
	return this
}

func (this QueryString) Limit2(offset int) QueryString {
	this.data = append(this.data, fmt.Sprintf(`limit(offset: %d)`, offset))
	return this
}
func (this QueryString) Limit3(limit int, offset int) QueryString {
	this.data = append(this.data, fmt.Sprintf(`limit(n: %d, offset: %d)`, limit, offset))
	return this
}

func (this QueryString) Bottom(n int) QueryString {
	this.data = append(this.data, fmt.Sprintf(`bottom(n:%d)`, n))
	return this
}

func (this QueryString) Columns(column string) QueryString {
	this.data = append(this.data, fmt.Sprintf(`columns(column: "%s")`, column))
	return this
}

func (this QueryString) First() QueryString {
	this.data = append(this.data, fmt.Sprintf(`first()`))
	return this
}

func (this QueryString) Count() QueryString {
	this.data = append(this.data, fmt.Sprintf(`count()`))
	return this
}

func (this QueryString) Last() QueryString {
	this.data = append(this.data, fmt.Sprintf(`last()`))
	return this
}

func (this QueryString) Find() string {
	return strings.Join(this.data, "\n\t|>\t")
}

func (this *Engine) Query(bucket string) QueryString {
	return QueryString{
		data: []string{fmt.Sprintf(`from(bucket:"%s")`, bucket)},
	}
}
