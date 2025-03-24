package main

import "testing"

func Test_isLeapYearJulian(t *testing.T) {
	tests := []struct {
		name string
		year int32
		want bool
	}{
		{
			name: "1804",
			year: 1804,
			want: true,
		},
		{
			name: "1805",
			year: 1805,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYearJulian(tt.year); got != tt.want {
				t.Errorf("isLeapYearJulian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLeapYearGregorian(t *testing.T) {
	tests := []struct {
		name string
		year int32
		want bool
	}{
		{
			name: "non-leap-1700",
			year: 1700,
			want: false,
		},
		{
			name: "non-leap-1800",
			year: 1800,
			want: false,
		},
		{
			name: "non-leap-1900",
			year: 1700,
			want: false,
		},
		{
			name: "non-leap-2100",
			year: 1700,
			want: false,
		},
		{
			name: "leap-1600",
			year: 1600,
			want: true,
		},
		{
			name: "leap-2000",
			year: 2000,
			want: true,
		},
		{
			name: "leap-2400",
			year: 1600,
			want: true,
		},
		{
			name: "leap-1944",
			year: 1944,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYearGregorian(tt.year); got != tt.want {
				t.Errorf("isLeapYearGregorian() = %v, want %v", got, tt.want)
			}
		})
	}
}
