package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var testCases = []struct {
	name string
	in   []int
	out  int
}{
	{
		name: "example 1",
		in:   []int{3, 4, -1, 1},
		out:  2,
	},
	{
		name: "big number",
		in:   []int{1, 2, 122},
		out:  3,
	},
	{
		name: "empty",
		in:   []int{},
		out:  1,
	},
	{
		name: "all there",
		in:   []int{1, 2, 3, 4},
		out:  5,
	},
	{
		name: "no such numbers",
		in:   []int{-1, -2, -3, -4},
		out:  1,
	},
	{
		name: "test",
		in:   []int{3, 2, 4, -1, 1},
		out:  5,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.in)
			if result != tc.out {
				t.Fatalf("expected %d, got %d", tc.out, result)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1000; i += 100 {
		slice := make([]int, i)
		for j := 0; j < len(slice); j++ {
			slice[j] = rand.Int()
		}
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solution(slice)
			}
		})
	}
}
