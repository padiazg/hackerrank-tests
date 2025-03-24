package main

import "testing"

func Test_catAndMouse(t *testing.T) {
	type args struct {
		x int32
		y int32
		z int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input",
			args: args{
				x: 2,
				y: 5,
				z: 4,
			},
			want: "Cat B",
		},
		{
			name: "input00-a",
			args: args{
				x: 1,
				y: 2,
				z: 3,
			},
			want: "Cat B",
		},
		{
			name: "input00-b",
			args: args{
				x: 1,
				y: 3,
				z: 2,
			},
			want: "Mouse C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catAndMouse(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("catAndMouse() = %v, want %v", got, tt.want)
			}
		})
	}
}
