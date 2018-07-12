package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(ss []string, p string) []string {
	tree := newTree()
	for _, s := range ss {
		tree.Load(s)
	}

	res := tree.Find(p)
	for i, r := range res {
		res[i] = p + r
	}
	return res
}

type node struct {
	value rune
	lefs  map[rune]*node
}

func newNode(v rune) *node {
	return &node{
		value: v,
		lefs:  map[rune]*node{},
	}
}

type tree struct {
	head *node
}

func newTree() *tree {
	return &tree{
		head: newNode('^'),
	}
}

func (t *tree) Find(s string) []string {
	n := t.head
	for _, c := range s {
		n = n.Find(c)
		if n == nil {
			return nil
		}
	}
	return n.String()
}

func (t *tree) Load(s string) {
	n := t.head
	for _, c := range s {
		n = n.Load(c)
	}
}

func (n *node) Find(r rune) *node {
	return n.lefs[r]
}

func (n *node) Load(r rune) *node {
	if found, ok := n.lefs[r]; ok {
		return found
	}
	new := newNode(r)
	n.lefs[r] = new
	return new
}

func (n *node) String() []string {
	r := []string{}
	for c, l := range n.lefs {
		if len(l.lefs) == 0 {
			r = append(r, string(c))
		}
		for _, s := range l.String() {
			r = append(r, string(c)+s)
		}
	}
	return r
}
