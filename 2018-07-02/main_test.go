package main

import "testing"

var testCases = []struct {
	name string
	a    []int
	k    int
	res  bool
}{
	{
		name: "From example",
		a:    []int{10, 15, 3, 7},
		k:    17,
		res:  true,
	},
	{
		name: "last digits",
		a:    []int{15, 3, 10, 7},
		k:    17,
		res:  true,
	},
	{
		name: "wrong result",
		a:    []int{15, 3, 10, 7},
		k:    1,
		res:  false,
	},
	{
		name: "empty array",
		a:    []int{},
		k:    12,
		res:  false,
	},
	{
		name: "1 element",
		a:    []int{1},
		k:    1,
		res:  false,
	},
	{
		name: "2 elements",
		a:    []int{1, 2},
		k:    3,
		res:  true,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem(tc.a, tc.k)
			if tc.res != result {
				t.Fatalf("expected %t, got %t", tc.res, result)
			}
		})
	}

}
