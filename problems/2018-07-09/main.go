package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

type node struct {
	value       int
	left, right *node
}

func unival(n *node) int {
	if n == nil {
		return 0
	}

	uni := true

	switch {
	case n.left != nil && n.right != nil:
		uni = n.left.value == n.right.value &&
			n.right.value == n.value
	case n.left != nil:
		uni = n.left.value == n.value
	case n.right != nil:
		uni = n.right.value == n.value
	}

	if !uni {
		return unival(n.left) + unival(n.right)
	}

	return unival(n.left) + unival(n.right) + 1
}
