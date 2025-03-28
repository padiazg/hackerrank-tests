package main

import "testing"

func Test_utopianTree(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{n: 5},
			want: 14,
		},
		{
			name: "input00",
			args: args{n: 1},
			want: 2,
		},
		{
			name: "input01",
			args: args{n: 4},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utopianTree(tt.args.n); got != tt.want {
				t.Errorf("utopianTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
