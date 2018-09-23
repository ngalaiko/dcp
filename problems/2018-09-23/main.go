package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(in string) string {
	out := strings.Builder{}
	currentChar := in[0]
	count := 0
	for _, c := range in {
		char := byte(c)
		switch char {
		case currentChar:
			count++
		default:
			out.WriteString(fmt.Sprintf("%d%c", count, currentChar))
			currentChar = char
			count = 1
		}
	}
	out.WriteString(fmt.Sprintf("%d%c", count, currentChar))
	return out.String()
}
