package main

import "testing"

func Test_hurdleRace(t *testing.T) {
	type args struct {
		k      int32
		height []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{
				k:      1,
				height: []int32{1, 2, 3, 3, 2},
			},
			want: 2,
		},
		{
			name: "input00",
			args: args{
				k:      4,
				height: []int32{1, 6, 3, 5, 2},
			},
			want: 2,
		},
		{
			name: "input01",
			args: args{
				k:      7,
				height: []int32{2, 5, 4, 5, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hurdleRace(tt.args.k, tt.args.height); got != tt.want {
				t.Errorf("hurdleRace() = %v, want %v", got, tt.want)
			}
		})
	}
}
