package tools

import (
	"sync"
	"time"
)

var syncMutex = sync.Mutex{}
var syncRWMutex = sync.Mutex{}

var syncMutexCache = NewCache[*sync.Mutex]()
var syncRWMutexCache = NewCache[*sync.RWMutex]()

func Mutex(key string) *sync.Mutex {
	syncMutex.Lock()
	defer syncMutex.Unlock()

	if o, ok := syncMutexCache.Get(key); ok {
		return o
	}

	oo := &sync.Mutex{}
	syncMutexCache.Set(key, oo, time.Now().Unix()+86400)
	return oo
}

func RWMutex(key string) *sync.RWMutex {
	syncRWMutex.Lock()
	defer syncRWMutex.Unlock()

	if o, ok := syncRWMutexCache.Get(key); ok {
		return o
	}

	oo := &sync.RWMutex{}
	syncRWMutexCache.Set(key, oo, time.Now().Unix()+86400)
	return oo
}
