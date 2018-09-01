package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		maze   [][]int
		start  []int
		finish []int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 int
	}{
		{
			name: "example",
			args: func(t *testing.T) args {
				return args{
					maze: [][]int{
						[]int{0, 0, 0, 0},
						[]int{-1, -1, 0, -1},
						[]int{0, 0, 0, 0},
						[]int{0, 0, 0, 0},
					},
					start:  []int{3, 0},
					finish: []int{0, 0},
				}
			},
			want1: 7,
		},
		{
			name: "long path",
			args: func(t *testing.T) args {
				return args{
					maze: [][]int{
						[]int{0, 0, 0, 0},
						[]int{-1, -1, -1, 0},
						[]int{0, 0, 0, 0},
						[]int{0, -1, -1, -1},
					},
					start:  []int{3, 0},
					finish: []int{0, 0},
				}
			},
			want1: 9,
		},
		{
			name: "no path",
			args: func(t *testing.T) args {
				return args{
					maze: [][]int{
						[]int{0, -1, 0, 0},
						[]int{-1, -1, -1, 0},
						[]int{0, 0, 0, 0},
						[]int{0, -1, -1, -1},
					},
					start:  []int{3, 0},
					finish: []int{0, 0},
				}
			},
			want1: 0,
		},
		{
			name: "2 paths",
			args: func(t *testing.T) args {
				return args{
					maze: [][]int{
						[]int{0, 0, 0, 0},
						[]int{0, -1, -1, 0},
						[]int{0, 0, 0, 0},
						[]int{0, -1, -1, -1},
					},
					start:  []int{3, 0},
					finish: []int{0, 0},
				}
			},
			want1: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.maze, tArgs.start, tArgs.finish)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
