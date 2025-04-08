package main

import (
	"io"
	"os"
	"testing"
)

func captureStdOut(f func()) string {
	var (
		orig    = os.Stdout
		r, w, _ = os.Pipe()
	)

	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}

func Test_matrixRotation(t *testing.T) {
	type args struct {
		matrix [][]int32
		r      int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input00",
			args: args{
				matrix: [][]int32{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
				r: 1,
			},
			want: `2 3 4 8
1 7 11 12
5 6 10 16
9 13 14 15`,
		},
		{
			name: "input11",
			args: args{
				matrix: [][]int32{
					{1, 2, 3, 4},
					{7, 8, 9, 10},
					{13, 14, 15, 16},
					{19, 2, 21, 22},
					{25, 26, 27, 28},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureStdOut(func() {
				matrixRotation(tt.args.matrix, tt.args.r)
			})
			if got != tt.want {
				t.Errorf("matrixRotation =\n%s\nwant \n%s", got, tt.want)
			}
		})
	}
}
