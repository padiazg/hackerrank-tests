package main

import "testing"

func Test_pickingNumbers(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{a: []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}},
			want: 5,
		},
		{
			args: args{a: []int32{4, 6, 5, 3, 3, 1}},
			want: 3,
		},
		{
			args: args{a: []int32{1, 2, 2, 3, 1, 2}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickingNumbers(tt.args.a); got != tt.want {
				t.Errorf("pickingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
