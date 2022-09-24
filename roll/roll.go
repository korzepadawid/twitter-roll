package roll

import "sync"

type Roll[T interface{}] struct {
	lock sync.RWMutex
	items []T
	capacity int
}

func New[T interface{}](cap int) *Roll[T] {
	return &Roll[T]{
		lock: sync.RWMutex{},
		items: make([]T, 0, cap),
		capacity: cap,
	}
}

func (r *Roll[T]) Add(item T) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if(len(r.items) < cap(r.items)){
		r.items = append([]T{item}, r.items...)
	} else {
		r.items = append([]T{item}, r.items[:len(r.items) - 1]...)
	}
}

func (r *Roll[T]) ReadAll() []T {
	r.lock.RLock()
	defer r.lock.RUnlock()
	res := r.items
	return res
}

func (r *Roll[T]) Flush() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.items = make([]T, 0, r.capacity)
}