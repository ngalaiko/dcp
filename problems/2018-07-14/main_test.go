package main

import "testing"

var testCases = []struct {
	name string
	inS  string
	inK  int
	out  string
}{
	{
		name: "example",
		inS:  "abcba",
		inK:  2,
		out:  "bcb",
	},
	{
		name: "3",
		inS:  "abcba",
		inK:  3,
		out:  "abcba",
	},
	{
		name: "4",
		inS:  "abcba",
		inK:  4,
		out:  "abcba",
	},
	{
		name: "5",
		inS:  "abcba",
		inK:  5,
		out:  "abcba",
	},
	{
		name: "0",
		inS:  "abcba",
		inK:  0,
		out:  "",
	},
	{
		name: "1",
		inS:  "abcba",
		inK:  1,
		out:  "a",
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.inK, tc.inS)
			if res != tc.out {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
