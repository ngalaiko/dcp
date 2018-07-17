package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

type Log struct {
	data  []string
	len   int
	write int
}

func NewLog(n int) *Log {
	return &Log{
		data: make([]string, 0, n),
	}
}

func (l *Log) Record(id string) {
	l.len++
	if len(l.data) < cap(l.data) {
		l.data = append(l.data, id)
		return
	}

	l.data[l.write] = id
	l.write++

	if l.write == cap(l.data) {
		l.write = 0
	}
}

func (l *Log) GetLast(i int) string {
	if i >= len(l.data) {
		return ""
	}

	return append(l.data[l.write:], l.data[:l.write]...)[len(l.data)-i-1]
}
