package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(aa []int) int {
	if len(aa) == 0 {
		return 1
	}
	for i := 0; i < len(aa); i++ {
		a := aa[i]
		if a < 0 { // we don't care, it's negative
			continue
		}
		if a >= len(aa) { // we don't care, result always < len(aa)
			continue
		}
		if a == aa[a-1] { // if integer on it's place, skip
			continue
		}
		// put each integer on it's place
		// decrease i, because aa[i] is a new integer now and we need to
		// process it one more time
		aa[i], aa[a-1] = aa[a-1], aa[i]
		i--
	}
	for i := range aa { // find first missing
		if aa[i] != i+1 {
			return i + 1
		}
	}
	return len(aa) + 1 // all there
}
