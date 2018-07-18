package main

import "testing"

var testCases = []struct {
	name string
	in   string
	out  int
}{
	{
		name: "example",
		in:   "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
		out:  20,
	},
	{
		name: "example 2",
		in:   "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
		out:  32,
	},
	{
		name: "file in first dir",
		in:   "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\t\t\tfile2.ext\n\tsubdir2\n\t\tsubsubdir2",
		out:  32,
	},
	{
		name: "no file",
		in:   "dir\n\tsubdir1\n\tsubdir2",
		out:  0,
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
