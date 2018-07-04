package main

import (
	"testing"
)

func Test(t *testing.T) {
	root := New(
		"root",
		New(
			"left",
			New(
				"left.left",
				nil,
				nil,
			),
			nil,
		),
		New(
			"right",
			nil,
			nil,
		),
	)

	rootS, err := root.Serialize()
	if err != nil {
		t.Fatalf("serialization error: %s", err)
	}

	node, err := Deserialize(rootS)
	if err != nil {
		t.Fatalf("deserialization error: %s", err)
	}

	if node.Left.Left.Value != "left.left" {
		t.Fatalf("test failed!")
	}
}
