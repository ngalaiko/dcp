package main

import (
	"fmt"
	"math/rand"
	"time"
)

// n - number of points inside square
// m - number of points inside circle
// pi = 4 * m / n
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	fmt.Println(Pi(10000000), time.Since(start))
	// 3.1420611999999988 1.306183355s
}

func Pi(n int) float64 {
	ng := 1000
	if n <= ng {
		return pi(float64(ng))
	}
	n = n / ng

	pis := make(chan float64, ng)
	for i := 0; i < ng; i++ {
		go func() {
			pis <- pi(float64(n))
		}()
	}

	var s float64
	var i int
	for p := range pis {
		s += p
		i++
		if i == ng {
			return s / float64(ng)
		}
	}
	return 0
}

func pi(n float64) float64 {
	var inCircle float64
	for i := 1.0; i < n; i++ {
		x, y := rand.Float64(), rand.Float64()

		insideCircle := (x-1)*(x-1)+(y-1)*(y-1) < 1
		if insideCircle {
			inCircle++
		}
	}
	return 4.0 * inCircle / n
}
