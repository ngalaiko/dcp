package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(s string) int {
	firstZero := s[0] == '0'
	l := len(s)
	switch {
	case l == 1:
		return canDecode(s)
	case l == 2:
		if firstZero {
			return canDecode(s)
		}
		return canDecode(s) + 1
	default:
		return canDecode(s[:1])*solution(s[1:]) +
			canDecode(s[:2])*solution(s[2:])
	}
}

// returns 1 if possible to decode string.
func canDecode(s string) int {
	if s[0] == '0' {
		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	if i >= 0 && i <= 26 {
		return 1
	}
	return 0
}
