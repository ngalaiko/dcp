package main

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

var re = regexp.MustCompile("\t")

func solution(p string) int {
	splitN := strings.Split(p, "\n")
	var max int
	for i := 0; i < len(splitN); i++ {
		if path.Ext(splitN[i]) == "" {
			continue
		}

		pathToFile := make([]string, i+1)
		for j := 0; j <= i; j++ {
			level := len(re.FindAllString(splitN[j], -1))
			pathToFile[level] = splitN[j]
		}
		var length int
		for _, s := range pathToFile {
			s = strings.Replace(s, "\t", "", -1)
			if len(s) == 0 {
				continue
			}
			length += len(s) + 1
		}
		if length-1 > max {
			max = length - 1
		}
	}
	return max
}
