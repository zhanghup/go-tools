package tools

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type ICache interface {
	Get(key string, result interface{}) bool
	Set(key string, obj interface{}, timeout ...int64)
	Delete(key string)
}

type cache struct {
	data sync.Map
}

// 是否自动清理内存中的过期数据
func CacheCreate(flag ...bool) ICache {
	c := &cache{data: sync.Map{}}

	if len(flag) > 0 && flag[0] {
		go func() {
			t := time.NewTicker(time.Second * 10)
			for {
				select {
				case <-t.C:
					keys := make([]interface{}, 0)
					c.data.Range(func(key, value interface{}) bool {
						dat2, ok := value.(cacheItem)
						if !ok {
							return true
						}
						if dat2.timeout != 0 && dat2.timeout <= time.Now().Unix() {
							keys = append(keys, key)
						}
						return true
					})
					for _, o := range keys {
						c.data.Delete(o)
					}
				}
			}
		}()
	}

	return c
}

type cacheItem struct {
	timeout int64
	data    interface{}
}

func (this *cache) Get(key string, result interface{}) bool {
	err := Ptr.Check(result)
	if err != nil {
		panic(errors.New("tools cache.Get result类型异常，应该为指针类型"))
	}
	data, ok := this.data.Load(key)
	if !ok {
		return ok
	}
	dat2, ok := data.(cacheItem)
	if !ok {
		return ok
	}
	if dat2.timeout == 0 || dat2.timeout > time.Now().Unix() {
		v1 := reflect.ValueOf(dat2.data)
		v2 := reflect.ValueOf(result)
		if v1.Kind() == reflect.Ptr {
			v2.Elem().Set(v1.Elem())
		} else {
			v2.Elem().Set(v1)
		}
		return true
	} else {
		this.data.Delete(key)
		return false
	}

}

func (this *cache) Set(key string, obj interface{}, timeout ...int64) {
	t := int64(0)
	if len(timeout) > 0 {
		t = timeout[0]
	}
	this.data.Store(key, cacheItem{t, obj})
}

func (this *cache) Delete(key string) {
	this.data.Delete(key)
}
