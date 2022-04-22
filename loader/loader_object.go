package loader

import (
	"sync"
	"time"
)

func NewObjectLoader[T any](fetch func(keys []string) (map[string]T, error)) IObject[T] {
	return &Object[T]{
		fetch:    fetch,
		wait:     time.Millisecond * 5,
		maxBatch: 500,
	}
}

type IObject[T any] interface {
	Load(key string) (T, bool, error)
}

// Object 批量缓存请求列表
type Object[T any] struct {
	fetch    func(keys []string) (map[string]T, error)
	wait     time.Duration
	maxBatch int
	batch    *objectLoaderBatch[T]
	mu       sync.Mutex
}

func (this *Object[T]) Wait(t time.Duration) {
	this.wait = t
}

func (this *Object[T]) MaxBatch(t int) {
	this.maxBatch = t
}

type objectLoaderBatch[T any] struct {
	keys    []string
	data    map[string]T
	error   error
	closing bool
	done    chan struct{}
}

func (l *Object[T]) Load(key string) (T, bool, error) {
	return l.LoadThunk(key)()
}

func (l *Object[T]) LoadThunk(key string) func() (T, bool, error) {

	l.mu.Lock()
	if l.batch == nil {
		l.batch = &objectLoaderBatch[T]{done: make(chan struct{})}
	}
	batch := l.batch
	batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (T, bool, error) {
		<-batch.done

		v, ok := batch.data[key]
		return v, ok, batch.error
	}
}

func (b *objectLoaderBatch[T]) keyIndex(l *Object[T], key string) {
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

func (b *objectLoaderBatch[T]) startTimer(l *Object[T]) {
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

func (b *objectLoaderBatch[T]) end(l *Object[T]) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
