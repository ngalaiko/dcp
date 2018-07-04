package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

type node struct {
	Value string `json:"value"`
	Left  *node  `json:"left"`
	Right *node  `json:"right"`
}

// New is a node constructor.
func New(v string, l, r *node) *node {
	return &node{
		Value: v,
		Left:  l,
		Right: r,
	}
}

// Serialize returns string representation of a node.
func (n *node) Serialize() (string, error) {
	b, err := json.Marshal(n)
	return string(b), err
}

// Deserialize returns node from it's string representation.
func Deserialize(s string) (*node, error) {
	n := new(node)
	return n, json.Unmarshal([]byte(s), n)
}
