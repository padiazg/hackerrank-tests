package main

import (
	"reflect"
	"testing"
)

func Test_circularArrayRotation(t *testing.T) {
	type args struct {
		a       []int32
		k       int32
		queries []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input",
			args: args{
				a:       []int32{3, 4, 5},
				k:       2,
				queries: []int32{1, 2},
			},
			want: []int32{5, 3},
		},
		{
			name: "input15",
			args: args{
				a:       []int32{1, 2, 3},
				k:       2,
				queries: []int32{0, 1, 2},
			},
			want: []int32{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayRotation(tt.args.a, tt.args.k, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularArrayRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
