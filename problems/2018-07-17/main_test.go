package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	for n := 1; n <= 10; n++ {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			l := NewLog(n)
			for i := 1; i <= 10; i++ {
				l.Record(strconv.Itoa(i))
			}

			for i := 0; i < n; i++ {
				res := l.GetLast(i)
				expected := strconv.Itoa(10 - i)
				if res != expected {
					t.Logf("expected %s, got %s", expected, res)
				}
			}
		})
	}
}
