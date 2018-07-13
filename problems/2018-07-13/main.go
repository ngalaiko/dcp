package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(n int, xx []int) int {
	if n == 0 {
		return 1
	}
	var res int
	for _, x := range xx {
		switch {
		case n == x:
			res += 1
		case n < x:
		default:
			res += solution(n-x, xx)
		}
	}
	return res
}
