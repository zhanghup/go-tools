package loader

import (
	"reflect"
	"sync"
	"time"
)

type ObjectFetch func(keys []string) (map[string]interface{}, error)

func NewObjectLoader(fetch ObjectFetch) IObject {
	return &Object{
		fetch:    fetch,
		wait:     time.Millisecond * 5,
		maxBatch: 500,
	}
}

type IObject interface {
	Load(key string, ru interface{}) (bool, error)
}

// Object 批量缓存请求列表
type Object struct {
	fetch ObjectFetch
	wait  time.Duration
	maxBatch int
	batch    *objectLoaderBatch
	mu       sync.Mutex
}

func (this *Object) Wait(t time.Duration) {
	this.wait = t
}

func (this *Object) MaxBatch(t int) {
	this.maxBatch = t
}

type objectLoaderBatch struct {
	keys    []string
	data    map[string]interface{}
	error   error
	closing bool
	done    chan struct{}
}

func (l *Object) Load(key string, ru interface{}) (bool, error) {
	res, err := l.LoadThunk(key)()
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	if ru == nil {
		return false, nil
	}

	vl := reflect.ValueOf(ru)

	if !vl.Elem().CanSet() {
		return false, nil
	}
	vl.Elem().Set(reflect.ValueOf(res))
	return true, nil
}

func (l *Object) LoadThunk(key string) func() (interface{}, error) {

	l.mu.Lock()
	if l.batch == nil {
		l.batch = &objectLoaderBatch{done: make(chan struct{})}
	}
	batch := l.batch
	batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (interface{}, error) {
		<-batch.done

		return batch.data[key], batch.error
	}
}

func (b *objectLoaderBatch) keyIndex(l *Object, key string) {
	for _, existingKey := range b.keys {
		if key == existingKey {
			return
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return
}

func (b *objectLoaderBatch) startTimer(l *Object) {
	time.Sleep(l.wait)
	l.mu.Lock()

	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *objectLoaderBatch) end(l *Object) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
