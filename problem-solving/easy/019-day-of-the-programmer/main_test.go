package main

import "testing"

func Test_dayOfProgrammer(t *testing.T) {
	tests := []struct {
		name string
		year int32
		want string
	}{
		{
			name: "input00",
			year: 2017,
			want: "13.09.2017",
		},
		{
			name: "input01",
			year: 2016,
			want: "12.09.2016",
		},
		{
			name: "input60",
			year: 1800,
			want: "12.09.1800",
		},
		{
			name: "input59",
			year: 1918,
			want: "26.09.1918",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfProgrammer(tt.year); got != tt.want {
				t.Errorf("dayOfProgrammer() = %v, want %v", got, tt.want)
			}
		})
	}
}
