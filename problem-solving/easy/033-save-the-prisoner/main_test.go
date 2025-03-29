package main

import "testing"

func Test_saveThePrisoner(t *testing.T) {
	type args struct {
		n int32
		m int32
		s int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{n: 4, m: 6, s: 2},
			want: 3,
		},
		{
			name: "input04a",
			args: args{n: 5, m: 2, s: 1},
			want: 2,
		},
		{
			name: "input04b",
			args: args{n: 5, m: 2, s: 2},
			want: 3,
		},
		{
			name: "input11a",
			args: args{n: 7, m: 19, s: 2},
			want: 6,
		},
		{
			name: "input11b",
			args: args{n: 3, m: 7, s: 3},
			want: 3,
		},
		{
			name: "input00-98",
			args: args{n: 499999999, m: 999999997, s: 2},
			want: 499999999,
		},
		{
			name: "input00-99",
			args: args{n: 499999999, m: 999999998, s: 2},
			want: 1,
		},
		{
			name: "input00-100",
			args: args{n: 999999999, m: 999999999, s: 1},
			want: 999999999,
		},
		{
			name: "input05-10",
			args: args{n: 13, m: 140874526, s: 1},
			want: 13,
		},
		{
			name: "input05-11",
			args: args{n: 5, m: 838370030, s: 1},
			want: 5,
		},
		{
			name: "input05-3",
			args: args{n: 2, m: 739424390, s: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := saveThePrisoner(tt.args.n, tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("saveThePrisoner() = %v, want %v", got, tt.want)
			}
		})
	}
}
