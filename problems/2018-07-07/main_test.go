package main

import (
	"testing"
)

func Test(t *testing.T) {
	l := list{}
	for i := 0; i < 10; i++ {
		l.Add(i)

		a, err := l.Get(i)
		if err != nil {
			t.Fatal(err)
		}

		if a != i {
			t.Fatalf("index %d: expected %d, got %d", i, i, a)
		}
	}

	_, err := l.Get(100)
	if err == nil {
		t.Fatal("should be error")
	}
}
