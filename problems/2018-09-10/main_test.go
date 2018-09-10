package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		reg string
		s   string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 bool
	}{
		{
			name:  ". test 1",
			args:  func(t *testing.T) args { return args{"ra.", "ray"} },
			want1: true,
		},
		{
			name:  ". test 2",
			args:  func(t *testing.T) args { return args{"ra.a", "raya"} },
			want1: true,
		},
		{
			name:  ". test 3",
			args:  func(t *testing.T) args { return args{".a.a", "raya"} },
			want1: true,
		},
		{
			name:  ". test 4",
			args:  func(t *testing.T) args { return args{"ra.", "raymond"} },
			want1: false,
		},
		{
			name:  ".* test 1",
			args:  func(t *testing.T) args { return args{".*at", "chat"} },
			want1: true,
		},
		{
			name:  ".* test 2",
			args:  func(t *testing.T) args { return args{".*", "chat"} },
			want1: true,
		},
		{
			name:  "*. test 3",
			args:  func(t *testing.T) args { return args{"*.", "chat"} },
			want1: true,
		},
		{
			name:  ".* test 4",
			args:  func(t *testing.T) args { return args{".*at", "chats"} },
			want1: false,
		},
		{
			name:  "* test 1",
			args:  func(t *testing.T) args { return args{"*", "chats"} },
			want1: true,
		},
		{
			name:  "* test 2",
			args:  func(t *testing.T) args { return args{"*", ""} },
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := solution(tArgs.reg, tArgs.s)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solution got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
