package main

import "testing"

func Test_getMoneySpent(t *testing.T) {
	type args struct {
		keyboards []int32
		drives    []int32
		b         int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{
				b:         60,
				keyboards: []int32{40, 50, 60},
				drives:    []int32{5, 8, 12},
			},
			want: 58,
		},
		{
			name: "input00",
			args: args{
				b:         10,
				keyboards: []int32{3, 1},
				drives:    []int32{5, 2, 8},
			},
			want: 9,
		},
		{
			name: "input01",
			args: args{
				b:         5,
				keyboards: []int32{4},
				drives:    []int32{5},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneySpent(tt.args.keyboards, tt.args.drives, tt.args.b); got != tt.want {
				t.Errorf("getMoneySpent() = %v, want %v", got, tt.want)
			}
		})
	}
}
