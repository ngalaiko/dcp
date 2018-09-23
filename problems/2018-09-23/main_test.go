package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "example",
			args: func(t *testing.T) args {
				return args{
					in: "AAAABBBCCDAA",
				}
			},
			want1: "4A3B2C1D2A",
		},
		{
			name: "one letter",
			args: func(t *testing.T) args {
				return args{
					in: "AAAAAAAA",
				}
			},
			want1: "8A",
		},
		{
			name: "different letters",
			args: func(t *testing.T) args {
				return args{
					in: "ABC",
				}
			},
			want1: "1A1B1C",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.in)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
