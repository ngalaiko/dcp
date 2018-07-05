package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(aa []int) int {
	if len(aa) == 0 {
		return 1
	}
	for _, a := range aa { // try to place all numbers at same index (from 1)
		if a < 0 { // we don't care, it's negative
			continue
		}
		if a >= len(aa) { // we don't care, result always < len(aa)
			continue
		}
		aa[a-1] = a
	}
	for i := range aa { // find first missing
		if aa[i] != i+1 {
			return i + 1
		}
	}
	return len(aa) + 1 // all there
}
