package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(ss string) bool {
	stack := &Stack{}
	for _, r := range ss {
		switch {
		case open[r]:
			stack.Push(r)
		case close[r]:
			top := stack.Pop()

			if top != reverse(r) {
				return false
			}
		}
	}
	return stack.Len() == 0
}

//
// helpers
//

// sets with open and closing brackets
var (
	open = map[rune]bool{
		'(': true,
		'{': true,
		'[': true,
	}
	close = map[rune]bool{
		')': true,
		'}': true,
		']': true,
	}
)

func reverse(r rune) rune {
	switch r {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	default:
		return ' '
	}
}

//
// Classic stack implementation on top of a slice.
//

type Stack struct {
	a []rune
}

func (s *Stack) Push(r rune) {
	s.a = append(s.a, r)
}

func (s *Stack) Pop() rune {
	end := len(s.a) - 1
	r := s.a[end]
	s.a = s.a[:end]
	return r
}

func (s *Stack) Len() int {
	return len(s.a)
}
