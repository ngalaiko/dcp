package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	res := <-solution(ch)

	fmt.Println(res)
}

func solution(in chan int) <-chan int {
	rand.Seed(time.Now().UTC().UnixNano())

	res := make(chan int, 1)

	go func() {
		var result int
		var lastprob float64
		for a := range in {
			prob := rand.Float64()
			if prob > lastprob {
				result = a
				lastprob = prob
			}
		}
		res <- result
	}()

	return res
}
