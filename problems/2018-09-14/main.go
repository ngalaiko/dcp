package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

type node struct {
	value int
	next  *node
}

func solution(k int, head *node) *node {
	// will contain tail of list with k+1 elements max
	// k+1 to be able to delete element later
	length := k + 1
	ptrs := make([]*node, length)

	c := head
	for c != nil {
		if len(ptrs) == length {
			ptrs = append(ptrs[1:], c)
		} else {
			ptrs = append(ptrs, c)
		}

		c = c.next
	}

	// at this point ptrs[1] is the element we need to delete
	if len(ptrs) < 3 {
		return head
	}

	ptrs[0].next = ptrs[2]

	return head
}

//
// helpers
//

func solutionWrapper(k int, aa ...int) []int {
	head := sliceToList(aa...)
	head = solution(k, head)
	return listToSlice(head)
}

func sliceToList(aa ...int) *node {
	head := &node{
		value: aa[0],
	}

	sliceToListHelper(head, aa[1:]...)

	return head
}

func sliceToListHelper(head *node, aa ...int) {
	if len(aa) == 0 {
		return
	}

	next := &node{
		value: aa[0],
	}
	head.next = next

	sliceToListHelper(next, aa[1:]...)
}

func listToSlice(head *node) []int {
	list := []int{}

	c := head
	for c != nil {
		list = append(list, c.value)
		c = c.next
	}

	return list
}
