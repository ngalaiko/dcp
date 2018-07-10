package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(aa []int) int {
	maxS := 0

	// 1st or 2nd should always be in sum
	for start := range make([]byte, 2) {
		// minimal step is 2
		for step := 2; step <= len(aa); step += 1 {
			s := 0
			for i := start; i < len(aa); i += step {
				s += aa[i]
			}

			if s > maxS {
				maxS = s
			}
		}
	}
	return maxS
}
