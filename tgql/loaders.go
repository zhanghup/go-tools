package tgql

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"net/http"
	"reflect"
	"sync"
	"time"
	"xorm.io/xorm"
)

type Loader interface {
	Object(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *ObjectLoader
	Slice(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *SliceLoader
}

const DATALOADEN_KEY = "go-app-dataloaden"

type dataLoaden struct {
	db *xorm.Engine

	objSync  sync.Locker
	objStore tools.ICache

	sliceSync  sync.Locker
	sliceStore tools.ICache
}

func NewDataLoaden(db *xorm.Engine) Loader {
	return &dataLoaden{
		db: db,

		objSync:  &sync.Mutex{},
		objStore: tools.CacheCreate(true),

		sliceSync:  &sync.Mutex{},
		sliceStore: tools.CacheCreate(true),
	}
}

func DataLoadenMiddleware(db *xorm.Engine, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), DATALOADEN_KEY, NewDataLoaden(db))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func DataLoaden(ctx context.Context) Loader {
	return ctx.Value(DATALOADEN_KEY).(Loader)
}

func (this *dataLoaden) Object(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *ObjectLoader {
	requestTable := reflect.TypeOf(table)
	if requestTable.Kind() == reflect.Ptr {
		requestTable = requestTable.Elem()
	}
	if requestTable.Kind() == reflect.Array || requestTable.Kind() == reflect.Slice {
		requestTable = requestTable.Elem()
		if requestTable.Kind() == reflect.Ptr {
			requestTable = requestTable.Elem()
		}
		if requestTable.Kind() != reflect.Struct {
			panic("table 数据结构异常")
		}
	}
	resultTable := requestTable
	if len(resultField) > 0 {
		f, ok := resultTable.FieldByName(resultField)
		if !ok {
			panic("resultField 不存在")
		}
		resultTable = f.Type
	}

	query := map[string]interface{}{}
	if param != nil {
		query = param
	}
	path := reflect.TypeOf(tools.Rft.RealValue(table)).PkgPath()
	name := reflect.TypeOf(tools.Rft.RealValue(table)).Name()
	key := fmt.Sprintf("table: %s/%s, sql: %s, param: %s, field: %s,resultField: %s", path, name, sql, tools.JSONString(query), keyField, resultField)
	key = tools.MD5([]byte(key))
	this.objSync.Lock()
	defer this.objSync.Unlock()
	if l := this.objStore.Get(key); l != nil {
		return l.(*ObjectLoader)
	}
	objLoader := &ObjectLoader{
		sync:         &sync.RWMutex{},
		db:           txorm.NewEngine(this.db),
		keyField:     keyField,
		resultField:  resultField,
		sql:          sql,
		param:        query,
		requestTable: requestTable,
		resultTable:  resultTable,
	}
	this.objStore.Set(key, objLoader, time.Now().Unix()+86400)
	return objLoader
}

func (this *dataLoaden) Slice(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *SliceLoader {
	requestTable := reflect.TypeOf(table)
	if requestTable.Kind() == reflect.Ptr {
		requestTable = requestTable.Elem()
	}
	if requestTable.Kind() == reflect.Array || requestTable.Kind() == reflect.Slice {
		requestTable = requestTable.Elem()
		if requestTable.Kind() == reflect.Ptr {
			requestTable = requestTable.Elem()
		}
		if requestTable.Kind() != reflect.Struct {
			panic("table 数据结构异常")
		}
	}
	resultTable := requestTable
	if len(resultField) > 0 {
		f, ok := resultTable.FieldByName(resultField)
		if !ok {
			panic("resultField 不存在")
		}
		resultTable = f.Type
	}

	query := map[string]interface{}{}
	if param != nil {
		query = param
	}

	path := reflect.TypeOf(tools.Rft.RealValue(table)).PkgPath()
	name := reflect.TypeOf(tools.Rft.RealValue(table)).Name()
	key := fmt.Sprintf("table: %s/%s, sql: %s, param: %s, field: %s,resultField: %s,resultTable: %s", path, name, sql, tools.JSONString(query), keyField, resultField, resultTable.PkgPath()+"/"+resultTable.Name())
	key = tools.MD5([]byte(key))

	this.sliceSync.Lock()
	defer this.sliceSync.Unlock()
	if l := this.sliceStore.Get(key); l != nil {
		return l.(*SliceLoader)
	}
	sliceLoader := &SliceLoader{
		sync:         &sync.RWMutex{},
		db:           txorm.NewEngine(this.db),
		keyField:     keyField,
		resultField:  resultField,
		sql:          sql,
		param:        query,
		requestTable: requestTable,
		resultTable:  resultTable,
	}
	this.sliceStore.Set(key, sliceLoader, time.Now().Unix()+86400)
	return sliceLoader
}
