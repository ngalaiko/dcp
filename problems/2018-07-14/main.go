package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(k int, s string) string {
	var res string
	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			sb := s[i:j]
			if numberOfDistinctChar(sb) > k {
				continue
			}
			if len(sb) > len(res) {
				res = sb
			}
		}
	}
	return res
}

func numberOfDistinctChar(s string) int {
	m := make(map[rune]struct{}, len(s))
	for _, c := range s {
		m[c] = struct{}{}
	}
	return len(m)
}
