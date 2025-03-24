package main

import "testing"

func Test_sockMerchant(t *testing.T) {
	type args struct {
		n  int32
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
				n:  7,
				ar: []int32{1, 2, 1, 2, 1, 3, 2},
			},
			want: 2,
		},
		{
			name: "input00",
			args: args{
				n:  7,
				ar: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			},
			want: 3,
		},
		{
			name: "input08",
			args: args{
				n:  7,
				ar: []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sockMerchant(tt.args.n, tt.args.ar); got != tt.want {
				t.Errorf("sockMerchant() = %v, want %v", got, tt.want)
			}
		})
	}
}
