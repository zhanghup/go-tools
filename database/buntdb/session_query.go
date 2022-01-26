package buntdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/buntdb"
	"reflect"
	"regexp"
	"strings"
)

type IQuery interface {
	Get(string, ...bool) (string, error)
	GetJson(string, interface{}, ...bool) error

	List(string, ListParam, func(string, string) bool) error
	ListJson(string, ListJsonParam, interface{}) error
}

type Query struct {
	tx *buntdb.Tx
}

func (this *Query) PageJson(index, size int, result interface{}) func(key, value string) (bool, error) {
	start := (index - 1) * size
	end := index * size

	count := 0

	typeArg := reflect.TypeOf(result).Elem().Elem()
	list := reflect.ValueOf(result).Elem()

	insert := func(val string) error {
		newValue := reflect.New(typeArg)
		err := json.Unmarshal([]byte(val), newValue.Interface())
		if err == nil {
			reflect.AppendSlice(list, newValue)
		}
		return err
	}

	return func(key, value string) (bool, error) {
		if size == -1 {
			return true, insert(value)
		}

		if count >= start && count < end {
			count += 1
			return true, insert(value)
		} else if count < start {
			count += 1
			return true, nil
		}

		return false, nil
	}
}

func (this *Query) JsonCheck(result interface{}) error {
	ty := reflect.TypeOf(result)
	if ty.Kind() != reflect.Ptr {
		return errors.New("result必须位指针值")
	}
	if ty.Elem().Kind() != reflect.Slice {
		return errors.New("指针必须指向一个数组")
	}
	return nil
}

type ListParamType string

type ListParam struct {
	Query string
	Order string // 排序方式 asc or desc
}

var (
	queryRegexpKey = regexp.MustCompile(`^key:\S+`)
	queryRegexpGt  = regexp.MustCompile(`^gt:\S+`)
	queryRegexpLt  = regexp.MustCompile(`^lt:\S+`)
	queryRegexpBt  = regexp.MustCompile(`^bt:\S+`)
)

/*
	List 列表查询，index 可以未空字符串

	ListParam.Query 查询参数为空则没有任何查询参数，返回索引内所有数据
	可选项：
		ListParam.Query = "任意非空字符串" 值匹配则返回
		ListParam.Query = "key:任意字符串" 值匹配则返回
		ListParam.Query = "gt:任意字符串" 若Order="asc" 则"值" >= "任意字符串"，若Order="desc" 则"值" > "任意字符串"
		ListParam.Query = "lt:任意字符串" 若Order="asc" 则"值" < "任意字符串"，若Order="desc" 则"值" <= "任意字符串"
		ListParam.Query = "bt:任意字符串,任意字符串" 若Order="asc" 则"任意字符串" <= "值" < "任意字符串"
		ListParam.Query = "bt:任意字符串,任意字符串" 若Order="desc" 则"任意字符串" < "值" <= "任意字符串"
*/
func (this *Query) List(index string, query ListParam, fn func(key, value string) bool) (err error) {
	queryString := strings.TrimSpace(query.Query)
	orderType := query.Order

	if len(queryString) > 0 {
		switch {
		case queryRegexpKey.MatchString(queryString):
			queryString = strings.Replace(queryString, "key:", "", 1)
			if orderType == "desc" {
				return this.tx.DescendKeys(queryString, fn)
			}
			return this.tx.AscendKeys(queryString, fn)
		case queryRegexpBt.MatchString(queryString):
			queryStrings := strings.Split(strings.Replace(queryString, "bt:", "", 1), ",")
			if len(queryStrings) == 2 {
				if orderType == "desc" {
					return this.tx.DescendRange(index, queryStrings[0], queryStrings[1], fn)
				}
				return this.tx.AscendRange(index, queryStrings[0], queryStrings[1], fn)
			} else {
				return errors.New("query格式错误")
			}
		case queryRegexpLt.MatchString(queryString):
			queryString = strings.Replace(queryString, "lt:", "", 1)
			if orderType == "desc" {
				return this.tx.DescendLessOrEqual(index, queryString, fn)
			}
			return this.tx.AscendLessThan(index, queryString, fn)
		case queryRegexpGt.MatchString(queryString):
			queryString = strings.Replace(queryString, "gt:", "", 1)
			if orderType == "desc" {
				return this.tx.DescendGreaterThan(index, queryString, fn)
			}
			return this.tx.AscendGreaterOrEqual(index, queryString, fn)
		default:
			if orderType == "desc" {
				return this.tx.DescendEqual(index, queryString, fn)
			}
			return this.tx.AscendEqual(index, queryString, fn)
		}
	} else {
		if orderType == "desc" {
			return this.tx.Descend(index, fn)
		}
		return this.tx.Ascend(index, fn)
	}
}

type ListJsonParam struct {
	ListParam

	Index int // 页码
	Size  int // 每页条数，-1返回所有
}

/*
	ListJson JSON列表查询，index 可以为空字符串
*/
func (this *Query) ListJson(index string, query ListJsonParam, result interface{}) (err error) {
	if err = this.JsonCheck(result); err != nil {
		return err
	}

	start := (query.Index - 1) * query.Size
	end := query.Index * query.Size

	count := 0
	items := make([]string, 0)
	var err2 error
	err = this.List(index, query.ListParam, func(key, value string) bool {
		if query.Size == -1 {
			items = append(items, value)
			return true
		}

		if count >= start && count < end {
			count += 1
			items = append(items, value)
			return true
		} else if count < start {
			count += 1
			return true
		}
		return false
	})
	if err2 != nil {
		return err2
	}
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), result)
}

func (this *Query) Get(key string, ignoreExpired ...bool) (val string, err error) {
	val, err = this.tx.Get(key, ignoreExpired...)
	if err == buntdb.ErrNotFound {
		return "", ErrNotFound
	}
	return
}

func (this *Query) GetJson(key string, result interface{}, ignoreExpired ...bool) error {
	val, err := this.Get(key, ignoreExpired...)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return err
	}
	return nil
}
