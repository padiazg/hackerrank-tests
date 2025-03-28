package main

import (
	"reflect"
	"testing"
)

func Test_climbingLeaderboard(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input00",
			args: args{
				ranked: []int32{100, 100, 50, 40, 40, 20, 10},
				player: []int32{5, 25, 50, 120},
			},
			want: []int32{6, 4, 2, 1},
		},
		{
			name: "input11",
			args: args{
				ranked: []int32{100, 90, 90, 80, 75, 60},
				player: []int32{50, 65, 77, 90, 102},
			},
			want: []int32{6, 5, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard(tt.args.ranked, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbingLeaderboard_a(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input00",
			args: args{
				ranked: []int32{100, 100, 50, 40, 40, 20, 10},
				player: []int32{5, 25, 50, 120},
			},
			want: []int32{6, 4, 2, 1},
		},
		{
			name: "input11",
			args: args{
				ranked: []int32{100, 90, 90, 80, 75, 60},
				player: []int32{50, 65, 77, 90, 102},
			},
			want: []int32{6, 5, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard_a(tt.args.ranked, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbingLeaderboard_b(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input00",
			args: args{
				ranked: []int32{100, 100, 50, 40, 40, 20, 10},
				player: []int32{5, 25, 50, 120},
			},
			want: []int32{6, 4, 2, 1},
		},
		{
			name: "input11",
			args: args{
				ranked: []int32{100, 90, 90, 80, 75, 60},
				player: []int32{50, 65, 77, 90, 102},
			},
			want: []int32{6, 5, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard_b(tt.args.ranked, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbingLeaderboard_c(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "input00",
			args: args{
				ranked: []int32{100, 100, 50, 40, 40, 20, 10},
				player: []int32{5, 25, 50, 120},
			},
			want: []int32{6, 4, 2, 1},
		},
		{
			name: "input11",
			args: args{
				ranked: []int32{100, 90, 90, 80, 75, 60},
				player: []int32{50, 65, 77, 90, 102},
			},
			want: []int32{6, 5, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard_c(tt.args.ranked, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}
