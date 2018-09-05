package main

import (
	"fmt"
	"strings"
)

func main() {
	root := New(1,
		New(2,
			New(4,
				nil,
				nil,
			),
			nil),
		New(3,
			nil,
			New(5,
				nil,
				nil,
			),
		),
	)

	fmt.Printf("initial: \n%s\n\n", root)
	fmt.Printf("root.left locked %t: \n%s\n\n", root.left.Lock(), root)
	fmt.Printf("root.left.left unlock %t: \n%s\n\n", root.left.left.Unlock(), root)
	fmt.Printf("root.left.left lock %t: \n%s\n\n", root.left.left.Lock(), root)
	fmt.Printf("root unlock %t: \n%s\n\n", root.Unlock(), root)
	fmt.Printf("root.left unlocked %t: \n%s\n\n", root.left.Unlock(), root)
	fmt.Printf("root lock %t: \n%s\n\n", root.Lock(), root)
	fmt.Printf("root.left lock %t: \n%s\n\n", root.left.Lock(), root)
	fmt.Printf("root.right.right lock %t: \n%s\n\n", root.right.right.Lock(), root)
}

type node struct {
	locked bool

	value       int
	parent      *node
	left, right *node
}

// New creates new node.
func New(v int, l, r *node) *node {
	n := &node{
		value: v,
		left:  l,
		right: r,
	}
	if l != nil {
		l.parent = n
	}
	if r != nil {
		r.parent = n
	}
	return n
}

// IsLocked returns true if node is locked.
func (n *node) IsLocked() bool {
	return n.locked
}

func (n *node) canChangeLock() bool {
	locked := n.parentLocked() ||
		n.childernLocked()
	return !locked
}

func (n *node) childernLocked() bool {
	locked := false
	if n.left != nil {
		locked = locked || n.left.locked || n.left.childernLocked()
	}
	if n.right != nil {
		locked = locked || n.right.locked || n.right.childernLocked()
	}
	return locked
}

func (n *node) parentLocked() bool {
	if n.parent != nil {
		return n.parent.locked || n.parent.parentLocked()
	}
	return false
}

// Lock locks the node, returns true if success.
func (n *node) Lock() bool {
	if !n.canChangeLock() {
		return false
	}
	n.locked = true
	return true
}

// Unlock unlocks the node, returns true if success.
func (n *node) Unlock() bool {
	if !n.canChangeLock() {
		return false
	}
	n.locked = false
	return true
}

// String implements Stringer.
func (n *node) String() string {
	return n.string(0)
}

func (n *node) string(t int) string {
	builder := &strings.Builder{}
	if t != 0 {
		builder.WriteString("\n")
	}
	for i := 0; i < t; i++ {
		builder.WriteString("	")
	}
	builder.WriteString(
		fmt.Sprintf("%d: %t", n.value, n.IsLocked()),
	)
	if n.left != nil {
		builder.WriteString(n.left.string(t + 1))
	}
	if n.right != nil {
		builder.WriteString(n.right.string(t + 1))
	}
	return builder.String()
}
