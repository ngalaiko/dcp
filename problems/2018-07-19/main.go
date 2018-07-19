package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(aa []int, k int) []int {
	res := make([]int, 0, len(aa)-k-1)
	for i := 0; i < len(aa)-k; i++ {
		m := 0
		for j := i; j <= i+k; j++ {
			if aa[j] > m {
				m = aa[j]
			}
		}
		res = append(res, m)
		fmt.Println(aa[i:i+k], m)
	}
	return res
}
