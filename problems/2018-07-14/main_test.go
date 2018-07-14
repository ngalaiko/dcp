package main

import "testing"

var testCases = []struct {
	name string
	in   interface{}
	out  interface{}
}{}

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
