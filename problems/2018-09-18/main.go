package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(words []string, k int) []string {
	lines := []string{}
	line := &strings.Builder{}
	for i := range words {
		word := words[i]
		switch {
		case line.Len() == 0: // first word in a line
			line.WriteString(word)
		case line.Len()+len(word)+1 <= k: // enough space for word + space
			line.WriteString(" ")
			line.WriteString(word)
		case line.Len() < k: // not enough space for a word, but there is some
			// insert spaces
			jLine := justifySpaces(line.String(), k)
			line = &strings.Builder{}
			line.WriteString(jLine)
			fallthrough
		default: // start next line
			lines = append(lines, line.String())
			line = &strings.Builder{}
			line.WriteString(word)
		}
	}
	return append(lines, justifySpaces(line.String(), k))
}

func justifySpaces(s string, k int) string {
	if len(s) == k {
		return s
	}
	if len(s) == 0 {
		return strings.Repeat(" ", k)
	}
	if !strings.Contains(s, " ") {
		s += " "
	}
	needSpaces := k - len(s)

	// extend spaces by 1 starting from left while length is not enough
	newString := &strings.Builder{}
	wasSpace := false
	space := ' '
	for _, r := range s {
		newString.WriteRune(r)

		if wasSpace && r == space {
			continue
		}

		wasSpace = false
		if r == space {
			wasSpace = true
			if needSpaces > 0 {
				newString.WriteRune(space)
				needSpaces--
			}
		}
	}
	return justifySpaces(newString.String(), k)
}
