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
		switch c {
		case currentChar:
			count++
		default:
			out.WriteString(fmt.Sprintf("%d%c"), count, currentChar)
			currentChar = c
			count = 1
		}
	}
	return out.String()
}
