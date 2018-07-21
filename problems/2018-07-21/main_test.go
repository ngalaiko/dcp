package main

import "testing"

var testCases = []struct {
	name string
	in1  *node
	in2  *node
	out  int
}{
	{
		name: "example",
		in1:  newNode(3, newNode(7, newNode(8, newNode(10, nil)))),
		in2:  newNode(99, newNode(1, newNode(8, newNode(10, nil)))),
		out:  8,
	},
	{
		name: "same",
		in1:  newNode(1, nil),
		in2:  newNode(1, nil),
		out:  1,
	},
	{
		name: "duplicate",
		in1:  newNode(8, newNode(10, newNode(8, newNode(10, nil)))),
		in2:  newNode(99, newNode(1, newNode(8, newNode(10, nil)))),
		out:  8,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.in1, tc.in2)
			if res != tc.out {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
