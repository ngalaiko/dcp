package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	name     string
	inString string
	inWords  []string
	out      [][]string
}{
	{
		name:     "example 1",
		inString: "thequickbrownfox",
		inWords:  []string{"quick", "brown", "the", "fox"},
		out: [][]string{
			[]string{"the", "quick", "brown", "fox"},
		},
	},
	{
		name:     "example 2",
		inString: "bedbathandbeyond",
		inWords:  []string{"bed", "bath", "bedbath", "and", "beyond"},
		out: [][]string{
			[]string{"bed", "bath", "and", "beyond"},
			[]string{"bedbath", "and", "beyond"},
		},
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.inString, tc.inWords)
			if fmt.Sprint(res) != fmt.Sprint(tc.out) {
				t.Fatalf("expected %v, got %v", tc.out, res)
			}
		})
	}
}
