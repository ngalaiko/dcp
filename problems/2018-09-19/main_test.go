package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 bool
	}{
		{
			name: "example 1",
			args: func(t *testing.T) args {
				return args{s: "([])[]({})"}
			},
			want1: true,
		},
		{
			name: "example 2",
			args: func(t *testing.T) args {
				return args{s: "([)]"}
			},
			want1: false,
		},
		{
			name: "example 3",
			args: func(t *testing.T) args {
				return args{s: "((()"}
			},
			want1: false,
		},
		{
			name: "empty strung",
			args: func(t *testing.T) args {
				return args{s: ""}
			},
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.s)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
