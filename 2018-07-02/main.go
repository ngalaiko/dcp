package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func problem(aa []int, k int) bool {
	for i, a := range aa {
		rest := k - a
		for j := i; j < len(aa); j++ {
			if aa[j] == rest {
				return true
			}
		}
	}

	return false
}
