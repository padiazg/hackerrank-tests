package main

import (
	"reflect"
	"testing"
)

func Test_absolutePermutation(t *testing.T) {
	type args struct {
		n int32
		k int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input",
			args: args{
				n: 4,
				k: 2,
			},
			want: []int32{3, 4, 1, 2},
		},
		{
			name: "input00a",
			args: args{
				n: 2,
				k: 1,
			},
			want: []int32{2, 1},
		},
		{
			name: "input00b",
			args: args{
				n: 3,
				k: 0,
			},
			want: []int32{1, 2, 3},
		},
		{
			name: "input00c",
			args: args{
				n: 3,
				k: 2,
			},
			want: []int32{-1},
		},
		{
			name: "input02b",
			args: args{
				n: 10,
				k: 1,
			},
			want: []int32{2, 1, 4, 3, 6, 5, 8, 7, 10, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absolutePermutation(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absolutePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
