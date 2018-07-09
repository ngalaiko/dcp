package main

import "testing"

var testCases = []struct {
	name string
}{}

func Test(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
