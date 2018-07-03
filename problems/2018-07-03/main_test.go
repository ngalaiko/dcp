package main

import "testing"

var testCases = []struct {
	name string
	in   []int
	out  []int
}{
	{
		name: "example 1",
		in:   []int{1, 2, 3, 4, 5},
		out:  []int{120, 60, 40, 30, 24},
	},
	{
		name: "example 2",
		in:   []int{3, 2, 1},
		out:  []int{2, 3, 6},
	},
	{
		name: "1 elem",
		in:   []int{5},
		out:  []int{5},
	},
	{
		name: "with 1 zero",
		in:   []int{3, 0, 1},
		out:  []int{0, 3, 0},
	},
	{
		name: "with 2 zero",
		in:   []int{3, 0, 1, 0},
		out:  []int{0, 0, 0, 0},
	},
}

func Test(t *testing.T) {
	test(t, solution)
	test(t, solution2)
}

func test(t *testing.T, s func([]int) []int) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			result := s(in)
			if len(result) != len(tc.out) {
				t.Logf("input: %v", tc.in)
				t.Logf("result: %v, expected: %v", result, tc.out)
				t.Fatalf("expected array of %d elements, got %d", len(result), len(tc.out))
			}
			for i := 0; i < len(tc.out); i++ {
				if tc.out[i] != result[i] {
					t.Logf("input: %v", tc.in)
					t.Logf("result: %v, expected: %v", result, tc.out)
					t.Fatalf("expected %d at index %d, got %d", tc.out[i], i, result[i])
				}
			}
		})
	}
}
