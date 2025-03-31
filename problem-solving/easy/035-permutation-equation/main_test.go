package main

import (
	"reflect"
	"testing"
)

func Test_permutationEquation(t *testing.T) {
	type args struct {
		p []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input",
			args: args{p: []int32{5, 2, 1, 3, 4}},
			want: []int32{4, 2, 5, 1, 3},
		},
		{
			name: "input00",
			args: args{p: []int32{2, 3, 1}},
			want: []int32{2, 3, 1},
		},
		{
			name: "input11",
			args: args{p: []int32{4, 3, 5, 1, 2}},
			want: []int32{1, 3, 5, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutationEquation(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutationEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
