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
		in:   []int{2, 4, 6, 2, 5},
		out:  13,
	},
	{
		name: "example 2",
		in:   []int{5, 1, 1, 5},
		out:  10,
	},
}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := solution(tc.in)
			if res != tc.out {
				t.Fatalf("expected %d, got %d", tc.out, res)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10000; i += 1000 {
		slice := make([]int, i)
		for j := 0; j < i; j++ {
			slice[j] = rand.Int()
		}
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solution(slice)
			}
		})
	}
}
