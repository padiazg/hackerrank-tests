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

func Test_bonAppetit(t *testing.T) {
	type args struct {
		bill []int32
		k    int32
		b    int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input",
			args: args{
				bill: []int32{2, 4, 6},
				k:    2,
				b:    6,
			},
			want: "3",
		},
		{
			name: "input",
			args: args{
				bill: []int32{2, 4, 6},
				k:    2,
				b:    3,
			},
			want: "Bon Appetit",
		},
		{
			name: "input00",
			args: args{
				bill: []int32{3, 10, 2, 9},
				k:    1,
				b:    12,
			},
			want: "5",
		},
		{
			name: "input06",
			args: args{
				bill: []int32{3, 10, 2, 9},
				k:    1,
				b:    7,
			},
			want: "Bon Appetit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureStdOut(func() {
				bonAppetit(tt.args.bill, tt.args.k, tt.args.b)
			})
			if got != tt.want {
				t.Errorf("captureStdOut = %s, want %s", got, tt.want)
			}
		})
	}
}
