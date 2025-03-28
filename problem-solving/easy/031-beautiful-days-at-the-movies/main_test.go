package main

import "testing"

func Test_beautifulDays(t *testing.T) {
	type args struct {
		i int32
		j int32
		k int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input08",
			args: args{i: 20, j: 23, k: 6},
			want: 2,
		},
		{
			name: "input09",
			args: args{i: 13, j: 45, k: 3},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulDays(tt.args.i, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("beautifulDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
