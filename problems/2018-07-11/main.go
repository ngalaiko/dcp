package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
}

func schedule(f func(), n int) {
	go func() {
		<-time.After(time.Duration(n) * time.Millisecond)
		f()
	}()
}
