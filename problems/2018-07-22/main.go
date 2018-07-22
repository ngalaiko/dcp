package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(tt [][]int) int {
	start := make(map[int]int, len(tt))
	finish := make(map[int]int, len(tt))
	from := tt[0][0]
	to := tt[0][1]
	for _, t := range tt {
		if t[0] < from {
			from = t[0]
		}
		if t[1] > to {
			to = t[1]
		}
		start[t[0]]++
		finish[t[1]]++
	}

	var all, buffer int
	for t := from; t <= to; t++ {
		if finishing, ok := finish[t]; ok {
			buffer += finishing
		}
		starting, ok := start[t]
		if !ok {
			continue
		}
		switch {
		case buffer >= starting:
			buffer -= starting
		case buffer >= 0 && buffer < starting:
			starting -= buffer
			buffer = 0
			fallthrough
		default:
			all += starting
		}
	}
	return all
}
