package tools

import (
	"reflect"
	"sync"
	"time"
)

type imap interface {
	Len() int
	Contain(k string) bool
	Get(k string) interface{}
	Set(k string, v interface{})
	Set2(k string, v interface{}, timeout int64)
	Remove(k string)
}

type safeobj struct {
	timeout int64
	obj     interface{}
}
type safeMap struct {
	Data map[string]safeobj
	*sync.RWMutex
}

func (d *safeMap) interval() {
	fn := func() {
		d.RLock()
		keys := reflect.ValueOf(d.Data).MapKeys()
		defer d.RUnlock()

		for _, k := range keys {
			go d.Contain(k.String())
		}
	}

	go func() {
		for {
			fn()
			time.Sleep(time.Minute)
		}
	}()
}
func (d *safeMap) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.Data)
}
func (d *safeMap) Contain(k string) bool {
	d.RLock()
	defer d.RUnlock()
	obj, ok := d.Data[k]
	if ok && obj.timeout != -1 && time.Now().Unix() > obj.timeout {
		go d.Remove(k)
		return false
	}

	return ok
}
func (d *safeMap) Get(k string) interface{} {
	if d.Contain(k) {
		d.RLock()
		defer d.RUnlock()
		return d.Data[k].obj
	}
	return nil
}
func (d *safeMap) Set(k string, v interface{}) {
	d.Lock()
	defer d.Unlock()
	d.Data[k] = safeobj{-1, v}
}
func (d *safeMap) Set2(k string, v interface{}, timeout int64) {
	d.Lock()
	defer d.Unlock()
	d.Data[k] = safeobj{timeout, v}
}
func (d *safeMap) Remove(k string) {
	d.Lock()
	defer d.Unlock()
	delete(d.Data, k)
}

func NewCache() imap {
	sfm := new(safeMap)
	sfm.Data = map[string]safeobj{}
	sfm.RWMutex = &sync.RWMutex{}
	sfm.interval()
	return sfm
}
