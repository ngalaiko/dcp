package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		k  int
		aa []int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 []int
	}{
		{
			name: "test1",
			args: func(*testing.T) args {
				return args{
					k:  3,
					aa: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				}
			},
			want1: []int{1, 2, 3, 4, 5, 6, 8, 9},
		},
		{
			name: "test2",
			args: func(*testing.T) args {
				return args{
					k:  7,
					aa: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				}
			},
			want1: []int{1, 2, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solutionWrapper(tArgs.k, tArgs.aa...)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
