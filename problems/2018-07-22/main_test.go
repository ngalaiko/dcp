package main

import "testing"

var testCases = []struct {
	name string
	in   [][]int
	out  int
}{
	{
		name: "example",
		in: [][]int{
			[]int{30, 75},
			[]int{0, 50},
			[]int{60, 150},
		},
		out: 2,
	},
	{
		name: "all same time",
		in: [][]int{
			[]int{10, 50},
			[]int{10, 50},
			[]int{10, 50},
		},
		out: 3,
	},
	{
		name: "one, then same time",
		in: [][]int{
			[]int{0, 5},
			[]int{10, 50},
			[]int{10, 50},
		},
		out: 2,
	},
	{
		name: "same time twice",
		in: [][]int{
			[]int{10, 50},
			[]int{10, 50},
			[]int{60, 90},
			[]int{60, 90},
		},
		out: 2,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.in)
			if res != tc.out {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
