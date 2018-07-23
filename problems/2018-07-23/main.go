package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(s string, ww []string) [][]string {
	sOrigin := s

	duplicatedWords := make(map[string]bool, len(ww))
	words := make([]string, 0, len(ww))
	for i := range ww {
		for j := i + 1; j < len(ww); j++ {
			if strings.Contains(ww[i], ww[j]) || strings.Contains(ww[j], ww[i]) {
				duplicatedWords[ww[i]] = true
				duplicatedWords[ww[j]] = true
			}
		}
		if !duplicatedWords[ww[i]] {
			words = append(words, ww[i])
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return strings.Index(s, words[i]) < strings.Index(s, words[j])
	})

	for _, word := range words {
		s = strings.Replace(s, word, "", -1)
	}

	if s == "" {
		return [][]string{words}
	}

	duplicatedWordsSlice := make([]string, 0, len(duplicatedWords))
	for word := range duplicatedWords {
		duplicatedWordsSlice = append(duplicatedWordsSlice, word)
	}

	result := make([][]string, 0, len(duplicatedWords))
	for i := range duplicatedWordsSlice {
		sCopy := s
		res := make([]string, 0, len(duplicatedWords))
		res = append(res, words...)
		for j := i; j < len(duplicatedWordsSlice); j++ {
			wasLen := len(sCopy)
			sCopy = strings.Replace(sCopy, duplicatedWordsSlice[j], "", -1)
			if wasLen == len(sCopy) {
				continue
			}
			res = append(res, duplicatedWordsSlice[j])
		}
		if len(sCopy) != 0 {
			continue
		}

		sort.Slice(res, func(i, j int) bool {
			return strings.Index(sOrigin, res[i]) < strings.Index(sOrigin, res[j])
		})
		result = append(result, res)
	}

	return result
}
