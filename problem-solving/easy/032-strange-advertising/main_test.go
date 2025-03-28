package main

import "testing"

func Test_viralAdvertising(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{n: 5},
			want: 24,
		},
		{
			name: "input00",
			args: args{n: 3},
			want: 9,
		},
		{
			name: "input01",
			args: args{n: 4},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viralAdvertising(tt.args.n); got != tt.want {
				t.Errorf("viralAdvertising() = %v, want %v", got, tt.want)
			}
		})
	}
}
