package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("Hello, world!")
}

type element struct {
	Value int
	both  uintptr
}

type list struct {
	head *element
}

func (l *list) Add(a int) {
	new := &element{
		Value: a,
	}
	newPtr := ptr(new)
	new.both = newPtr

	if l.head == nil {
		l.head = new
		return
	}

	var tail, previous *element
	l.Range(func(e *element) bool {
		tail = e
		previous = tail
		return true
	})

	tail.both = uintptr(unsafe.Pointer(ptr(previous) ^ newPtr))
}

func (l *list) Range(f func(*element) bool) {
	if l.head == nil {
		return
	}

	h := l.head

	var previous *element
	for {
		if !f(h) {
			return
		}

		headPtr := ptr(h)
		if headPtr == h.both {
			return
		}

		previous = h
		previousPtr := ptr(previous)

		h = (*element)(unsafe.Pointer(h.both ^ previousPtr))
	}
}

func (l *list) Get(i int) (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("out of range")
	}
	if i < 0 {
		return 0, fmt.Errorf("should be positive")
	}

	k := 0
	var res int
	l.Range(func(e *element) bool {
		res = e.Value
		if k == i {
			return false
		}
		k++
		return true
	})

	if i != k {
		return 0, fmt.Errorf("out of range")
	}

	return res, nil
}

func ptr(e *element) uintptr {
	return uintptr(unsafe.Pointer(e))
}
