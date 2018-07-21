package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(a, b *node) int {
	a = flip(a)
	b = flip(b)
	var res int
	for a != nil && b != nil && a.value == b.value {
		res = a.value
		a = a.next
		b = b.next
	}
	return res
}

func flip(n *node) *node {
	var c *node = nil
	for n != nil {
		next := n.next
		n.next = c

		c = n
		n = next
	}
	return c
}

type node struct {
	value int
	next  *node
}

func newNode(v int, next *node) *node {
	return &node{
		value: v,
		next:  next,
	}
}
