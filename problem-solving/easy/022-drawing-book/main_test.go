package main

import "testing"

func Test_pageCount(t *testing.T) {
	type args struct {
		n int32
		p int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "first-page",
			args: args{
				n: 5,
				p: 1,
			},
			want: 0,
		},
		{
			name: "last-page-non-pair",
			args: args{
				n: 5,
				p: 5,
			},
			want: 0,
		},
		{
			name: "prior-last-page-non-pair",
			args: args{
				n: 5,
				p: 4,
			},
			want: 0,
		},
		{
			name: "input",
			args: args{
				n: 5,
				p: 3,
			},
			want: 1,
		},
		{
			name: "input00",
			args: args{
				n: 6,
				p: 2,
			},
			want: 1,
		},
		{
			name: "input01",
			args: args{
				n: 5,
				p: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageCount(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("pageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
