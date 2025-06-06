package main

import "testing"

func Test_divisibleSumPairs(t *testing.T) {
	type args struct {
		k  int32
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{
				k:  5,
				ar: []int32{1, 2, 3, 4, 5, 6},
			},
			want: 3,
		},
		{
			name: "input00",
			args: args{
				k:  3,
				ar: []int32{1, 3, 2, 6, 1, 2},
			},
			want: 5,
		},
		{
			name: "input12",
			args: args{
				k:  77,
				ar: []int32{85, 42, 54, 62, 79, 71, 29, 61, 1, 92, 93, 99, 82, 5, 45, 55, 49, 49, 93, 45, 2, 53, 80, 68, 51, 15, 42, 8, 5, 45, 95, 90, 4, 5, 45, 56, 20, 66, 32, 65, 11, 48, 41, 10, 92, 41, 8, 23, 88, 50, 90, 2, 3, 88, 29, 34, 54, 83, 91, 37, 95, 11, 7, 48, 96, 2, 84, 50, 20, 97, 95, 85, 80, 87, 99, 34, 40, 33, 78, 6, 58, 82, 49, 37, 35, 12, 85, 73, 96, 7, 63, 36, 73, 3, 96, 23, 5, 75, 16, 41},
			},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := int32(len(tt.args.ar))
			if got := divisibleSumPairs(n, tt.args.k, tt.args.ar); got != tt.want {
				t.Errorf("divisibleSumPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
