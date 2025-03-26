package main

import "testing"

func Test_formingMagicSquare(t *testing.T) {
	type args struct {
		s [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{s: [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}},
			want: 7,
		},
		{
			name: "input",
			args: args{s: [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}},

			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formingMagicSquare(tt.args.s); got != tt.want {
				t.Errorf("formingMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
