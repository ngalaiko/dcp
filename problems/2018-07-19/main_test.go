package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	name string
	inN  int
	inX  []int
	out  []int
}{
	{
		name: "example",
		inN:  2,
		inX:  []int{10, 5, 2, 7, 8, 7},
		out:  []int{10, 7, 8, 8},
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.inX, tc.inN)
			if fmt.Sprint(res) != fmt.Sprint(tc.out) {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
