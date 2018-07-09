package main

import "testing"

var testCases = []struct {
	in  string
	out int
}{
	{"1234", 3},
	{"111", 3},
	{"1", 1},
	{"14", 2},
	{"143", 2},
	{"276", 1},
	{"1212", 5},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			res := solution(tc.in)
			if res != tc.out {
				t.Fatalf("expected %d, got %d", tc.out, res)
			}
		})
	}
}
