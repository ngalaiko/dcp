package main

import "testing"

var testCases = []struct {
	in  *node
	out int
}{
	{
		in: &node{
			value: 0,
			left:  &node{value: 1},
			right: &node{
				value: 0,
				left: &node{
					value: 1,
					left:  &node{value: 1},
					right: &node{value: 1},
				},
				right: &node{value: 0},
			},
		},
		out: 5,
	},
	{
		in: &node{
			value: 1,
		},
		out: 1,
	},
	{
		in: &node{
			value: 1,
			left: &node{
				value: 0,
			},
		},
		out: 1,
	},
	{
		in: &node{
			value: 1,
			left: &node{
				value: 0,
			},
			right: &node{
				value: 1,
			},
		},
		out: 2,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		res := unival(tc.in)
		if res != tc.out {
			t.Fatalf("expected %d, got %d", tc.out, res)
		}
	}
}
