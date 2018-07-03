package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

func solution(aa []int) []int {
	if len(aa) < 2 {
		return aa
	}
	result := 1
	countZero := 0
	for _, a := range aa {
		if a == 0 {
			countZero++
			continue
		}
		result *= a
	}
	if countZero > 1 {
		return make([]int, len(aa))
	}
	for i := range aa {
		if aa[i] == 0 {
			aa[i] = result
			continue
		}
		if countZero == 1 {
			aa[i] = 0
			continue
		}
		aa[i] = result / aa[i]
	}
	return aa
}

func solution2(aa []int) []int {
	if len(aa) < 2 {
		return aa
	}
	result := make([]int, len(aa))
	for i := range aa {
		v := 1
		for j, aj := range aa {
			if i == j {
				continue
			}
			v *= aj
		}
		result[i] = v
	}
	return result
}
