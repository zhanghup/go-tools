package tools

import (
	"sync"
	"time"
)

var syncMutex = sync.Mutex{}
var syncRWMutex = sync.Mutex{}
var syncMutexCache = CacheCreate(true)
var syncRWMutexCache = CacheCreate(true)

func Mutex(key string) *sync.Mutex {
	syncMutex.Lock()
	defer syncMutex.Unlock()
	if o := syncMutexCache.Get(key); o != nil {
		return o.(*sync.Mutex)
	}

	oo := &sync.Mutex{}
	syncMutexCache.Set(key, oo, time.Now().Unix()+86400)
	return oo
}

func RWMutex(key string) *sync.RWMutex {
	syncRWMutex.Lock()
	defer syncRWMutex.Unlock()

	if o := syncRWMutexCache.Get(key); o != nil {
		return o.(*sync.RWMutex)
	}

	oo := &sync.RWMutex{}
	syncRWMutexCache.Set(key, oo, time.Now().Unix()+86400)
	return oo
}
