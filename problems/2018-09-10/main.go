package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(reg string, s string) bool {
	currentIndex := 0
	var i int
	for i = 0; i < len(reg); i++ {
		switch reg[i] {
		case '.':
			// skip 1 char
			currentIndex++
		case '*':
			if i == len(reg)-1 {
				// in case * is the last symbol in regex, it is valid anyway
				return true
			}
			// skip chars until we have same number of
			// chars in regex after *
			// and string after current index
			until := len(s) - (len(reg) - i - 1)
			for currentIndex < until && s[currentIndex] != reg[i+1] {
				currentIndex++
			}
		default:
			// check if valid char at position
			if reg[i] != s[currentIndex] {
				return false
			}
			currentIndex++
		}
	}

	// if string was completely validated unsing complete regex, return true
	return currentIndex == len(s) && i == len(reg)
}
