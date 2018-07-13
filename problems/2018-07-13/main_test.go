package main

import "testing"

var testCases = []struct {
	name string
	inN  int
	inX  []int
	out  int
}{
	{
		name: "example",
		inN:  4,
		inX:  []int{1, 2},
		out:  5,
	},
	{
		name: "1 way",
		inN:  4,
		inX:  []int{1},
		out:  1,
	},
	{
		name: "2 way",
		inN:  4,
		inX:  []int{1, 4},
		out:  2,
	},
	{
		name: "0 way",
		inN:  4,
		inX:  []int{},
		out:  0,
	},
	{
		name: "1 more way",
		inN:  4,
		inX:  []int{4, 2, 1},
		out:  6,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.inN, tc.inX)
			if res != tc.out {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
