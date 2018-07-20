package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(mm [][]int) int {
	n, m := len(mm), len(mm[0])
	if m > n {
		n, m = m, n
	}

	rowSum1 := n / 2 * 3 // sum of odd row
	if n%2 == 0 {
		return rowSum1 * m
	}

	return m/2*(rowSum1+1) + m/2*(rowSum1+2) + m%2*(rowSum1+1)
}
