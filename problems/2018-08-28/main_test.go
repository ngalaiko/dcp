package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		maze   []int
		_      int
		start  []int
		finish []int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 int
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.maze, tArgs._, tArgs.start, tArgs.finish)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
