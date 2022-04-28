package tools

import "sync"

type Queue[T any] struct {
	data []T

	lock sync.Mutex
}

type IQueue[T any] interface {
	Push(ele ...T)
	LPush(ele ...T)
	Pop(n int) []T
	Len() int
}

func NewQueue[T any]() IQueue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (this *Queue[T]) Push(ele ...T) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.data = append(this.data, ele...)
}

func (this *Queue[T]) LPush(ele ...T) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.data = append(ele, this.data...)
}

func (this *Queue[T]) Pop(n int) []T {
	this.lock.Lock()
	defer this.lock.Unlock()

	if n > len(this.data) {
		n = len(this.data)
	}

	data := this.data[:n]
	this.data = this.data[n:]
	return data
}

func (this *Queue[T]) Len() int {
	return len(this.data)
}
