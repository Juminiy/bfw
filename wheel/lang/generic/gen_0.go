package generic

import "sync"

type LockInterface[T any] interface {
	Read() T
	Write(T) T
}

type Lockable[T any] struct {
	Elem T
	sync.Mutex
}

func (l *Lockable[T]) Read() T {
	var elem T
	l.Lock()
	elem = l.Elem
	l.Unlock()
	return elem
}

func (l *Lockable[T]) Write(elem T) T {
	var preElem T
	l.Lock()
	preElem = l.Elem
	l.Elem = elem
	l.Unlock()
	return preElem
}

type RWLockable[T any] struct {
	Elem T
	sync.RWMutex
}

func (l *RWLockable[T]) Read() T {
	var elem T
	l.RLock()
	elem = l.Elem
	l.RUnlock()
	return elem
}

func (l *RWLockable[T]) Write(elem T) T {
	var preElem T
	l.Lock()
	preElem = l.Elem
	l.Elem = elem
	l.Unlock()
	return preElem
}
