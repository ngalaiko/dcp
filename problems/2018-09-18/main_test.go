package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		ss []string
		k  int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 []string
	}{
		{
			name: "example",
			args: func(t *testing.T) args {
				return args{
					ss: []string{"the", "quick", "brown", "fox", "jumps",
						"over", "the", "lazy", "dog"},
					k: 16,
				}
			},
			want1: []string{
				"the  quick brown",
				"fox  jumps  over",
				"the   lazy   dog",
			},
		},
		{
			name: "all spaces",
			args: func(t *testing.T) args {
				return args{
					ss: []string{""},
					k:  16,
				}
			},
			want1: []string{
				"                ",
			},
		},
		{
			name: "add spaces",
			args: func(t *testing.T) args {
				return args{
					ss: []string{"abc", "da"},
					k:  5,
				}
			},
			want1: []string{
				"abc  ",
				"da   ",
			},
		},
		{
			name: "concat",
			args: func(t *testing.T) args {
				return args{
					ss: []string{"ac", "da"},
					k:  5,
				}
			},
			want1: []string{
				"ac da",
			},
		},
		{
			name: "do not change",
			args: func(t *testing.T) args {
				return args{
					ss: []string{"ac", "da"},
					k:  2,
				}
			},
			want1: []string{
				"ac",
				"da",
			},
		},
		{
			name: "string longer then k",
			args: func(t *testing.T) args {
				return args{
					ss: []string{"ac", "sas", "da"},
					k:  2,
				}
			},
			want1: []string{
				"ac",
				"sas",
				"da",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.ss, tArgs.k)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
