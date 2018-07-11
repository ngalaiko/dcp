package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	res := []int{}
	rand.Seed(time.Now().UTC().UnixNano())
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		res = append(res)

		v := rand.Intn(100) * 10

		wg.Add(1)
		schedule(func() {
			res = append(res, v)
			wg.Done()
		}, v)
	}

	wg.Wait()
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			t.Log(res)
			t.Fatalf("index %d. %d > %d", i, res[i], res[i+1])
		}
	}
}
