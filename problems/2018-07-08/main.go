package main

import (
	"fmt"
	"unsafe"
)

func main() {
	l := list{}
	for i := 0; i < 10; i++ {
		l.Add(i)
		fmt.Println("added")

		a, err := l.Get(i)
		if err != nil {
			fmt.Printf("%d: %s\n", i, err)
			continue
		}

		fmt.Printf("%d: %d\n", i, a)
	}

	_, err := l.Get(100)
	fmt.Printf("%d: %s\n", 100, err)
}

type element struct {
	value int
	both  uintptr
}

type list struct {
	head *element
}

func (l *list) Add(a int) {
	new := &element{
		value: a,
	}
	newPtr := ptr(new)
	new.both = newPtr

	if l.head == nil {
		l.head = new
		return
	}

	h := l.head
	headPtr := ptr(h)
	for headPtr != h.both {
		h = (*element)(unsafe.Pointer(headPtr | h.both))
	}

	h.both = uintptr(unsafe.Pointer(headPtr ^ newPtr))
}

func (l *list) Get(i int) (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("out of range")
	}
	if i < 0 {
		return 0, fmt.Errorf("should be positive")
	}

	k := 0
	h := l.head
	headPtr := ptr(h)
	fmt.Println(headPtr, h.both)
	for headPtr != h.both && i != k {
		h = (*element)(unsafe.Pointer(headPtr | h.both))
	}

	if i != k {
		return 0, fmt.Errorf("out of range")
	}

	return h.value, nil
}

func ptr(e *element) uintptr {
	return uintptr(unsafe.Pointer(e))
}
