package main

import "testing"

var testCases = []struct {
	name string
	in   [][]int
	out  int
}{
	{
		name: "1x1",
		in: [][]int{
			[]int{1},
		},
		out: 1,
	},
	{
		name: "1x2",
		in: [][]int{
			[]int{1, 2},
		},
		out: 3,
	},
	{
		name: "1x3",
		in: [][]int{
			[]int{1, 2, 1},
		},
		out: 4,
	},
	{
		name: "1x4",
		in: [][]int{
			[]int{1, 2, 1, 2},
		},
		out: 6,
	},
	{
		name: "1x5",
		in: [][]int{
			[]int{1, 2, 1, 2, 1},
		},
		out: 7,
	},
	{
		name: "1x6",
		in: [][]int{
			[]int{1, 2, 1, 2, 1, 2},
		},
		out: 9,
	},
	{
		name: "1x7",
		in: [][]int{
			[]int{1, 2, 1, 2, 1, 2, 1},
		},
		out: 10,
	},
	{
		name: "2x2",
		in: [][]int{
			[]int{1, 2},
			[]int{2, 1},
		},
		out: 6,
	},
	{
		name: "2x3",
		in: [][]int{
			[]int{1, 2, 1},
			[]int{2, 1, 2},
		},
		out: 9,
	},
	{
		name: "3x3",
		in: [][]int{
			[]int{1, 2, 1},
			[]int{2, 1, 2},
			[]int{1, 2, 1},
		},
		out: 13,
	},
	{
		name: "2x4",
		in: [][]int{
			[]int{1, 2, 1, 2},
			[]int{2, 1, 2, 1},
		},
		out: 12,
	},
	{
		name: "3x4",
		in: [][]int{
			[]int{1, 2, 1, 2},
			[]int{2, 1, 2, 1},
			[]int{1, 2, 1, 2},
		},
		out: 18,
	},
	{
		name: "4x4",
		in: [][]int{
			[]int{1, 2, 1, 2},
			[]int{2, 1, 2, 1},
			[]int{1, 2, 1, 2},
			[]int{2, 1, 2, 1},
		},
		out: 24,
	},
	{
		name: "7x6",
		in: [][]int{
			[]int{1, 2, 1, 2, 1, 2},
			[]int{2, 1, 2, 1, 2, 1},
			[]int{1, 2, 1, 2, 1, 2},
			[]int{2, 1, 2, 1, 2, 1},
			[]int{1, 2, 1, 2, 1, 2},
			[]int{2, 1, 2, 1, 2, 1},
			[]int{1, 2, 1, 2, 1, 2},
		},
		out: 63,
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
