package main

import "testing"

var testCases = []struct {
	name string
	in2  string
	in   []string
	out  []string
}{
	{
		name: "example 1",
		in:   []string{"dog", "deer", "deal"},
		in2:  "de",
		out:  []string{"deer", "deal"},
	},
	{
		name: "no suggestion",
		in:   []string{"dog", "deer", "deal"},
		in2:  "ga",
		out:  []string{},
	},
	{
		name: "two roots",
		in:   []string{"dog", "game", "grade"},
		in2:  "g",
		out:  []string{"game", "grade"},
	},
	{
		name: "full word",
		in:   []string{"testing", "test", "a"},
		in2:  "testing",
		out:  []string{},
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.in, tc.in2)
			if len(res) != len(tc.out) {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}

			resSet := map[string]bool{}
			for _, s := range tc.out {
				resSet[s] = true
			}
			for _, r := range res {
				if !resSet[r] {
					t.Fatalf("expected %v, got %v", tc.out, res)
				}
			}
		})
	}
}
