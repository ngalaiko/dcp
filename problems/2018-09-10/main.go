package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(reg string, s string) bool {
	for _, rc := range reg {
		switch rc {
		case '.':
		case '*':
		default:
		}
	}

	return false
}
