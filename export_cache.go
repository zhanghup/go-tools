package tools

import (
	"sync"
	"time"
)

type ICache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, obj T, timeout ...int64)
	Delete(key string)
	Exist(key string) bool
	Clear()
	Data() map[string]T
	Len() int
}

type cache[T any] struct {
	data map[string]cacheItem[T]
	rw   sync.RWMutex
	once sync.Once
}

type cacheItem[T any] struct {
	timeout int64
	data    T
}

func NewCache[T any]() ICache[T] {
	c := &cache[T]{data: map[string]cacheItem[T]{}}
	c.onclear()
	return c
}

func (this *cache[T]) onclear() {
	go this.once.Do(func() {
		t := time.NewTicker(time.Second * 30) // 30秒清理一次
		for {
			select {
			case <-t.C:
				Run(this.Clear)
			}
		}
	})
}

func (this *cache[T]) Exist(key string) bool {
	_, ok := this.Get(key)
	return ok
}

func (this *cache[T]) Get(key string) (T, bool) {
	this.rw.RLock()
	defer this.rw.RUnlock()

	o, ok := this.data[key]
	if !ok {
		return o.data, false
	}
	if o.timeout == 0 {
		return o.data, true
	}
	if o.timeout > time.Now().Unix() {
		return o.data, true
	}
	return o.data, false
}

// Set timeout 具体时间戳
func (this *cache[T]) Set(key string, obj T, timeout ...int64) {
	this.rw.Lock()
	defer this.rw.Unlock()

	if len(timeout) > 0 {
		this.data[key] = cacheItem[T]{
			timeout: timeout[0],
			data:    obj,
		}
	} else {
		this.data[key] = cacheItem[T]{
			timeout: 0,
			data:    obj,
		}
	}
}

func (this *cache[T]) Delete(key string) {
	this.rw.Lock()
	defer this.rw.Unlock()

	delete(this.data, key)
}

func (this *cache[T]) Data() map[string]T {
	this.rw.RLock()
	defer this.rw.RUnlock()
	now := time.Now().Unix()
	result := map[string]T{}

	for k, v := range this.data {
		if v.timeout == 0 || v.timeout > now {
			result[k] = v.data
		}
	}
	return result
}

func (this *cache[T]) Len() int {
	this.rw.RLock()
	defer this.rw.RUnlock()
	now := time.Now().Unix()

	n := 0

	for _, v := range this.data {
		if v.timeout == 0 || v.timeout > now {
			n += 1
		}
	}
	return n
}

func (this *cache[T]) Clear() {
	this.rw.Lock()
	defer this.rw.Unlock()

	now := time.Now().Unix()
	keys := make([]string, 0)

	for k, v := range this.data {
		if v.timeout != 0 && v.timeout < now {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(this.data, k)
	}
}
