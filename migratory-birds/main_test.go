package main

import "testing"

func Test_migratoryBirds(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{
				arr: []int32{1, 1, 2, 2, 3},
			},
			want: 1,
		},
		{
			name: "input00",
			args: args{
				arr: []int32{1, 4, 4, 4, 5, 3},
			},
			want: 4,
		},
		{
			name: "input05",
			args: args{
				arr: []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := migratoryBirds(tt.args.arr); got != tt.want {
				t.Errorf("migratoryBirds() = %v, want %v", got, tt.want)
			}
		})
	}
}
