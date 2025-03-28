package main

import "testing"

func Test_angryProfessor(t *testing.T) {
	type args struct {
		k int32
		a []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input",
			args: args{
				k: 3,
				a: []int32{-2, -1, 0, 1, 2},
			},
			want: "NO",
		},
		{
			name: "input00a",
			args: args{
				k: 3,
				a: []int32{-1, -3, 4, 2},
			},
			want: "YES",
		},
		{
			name: "input00b",
			args: args{
				k: 2,
				a: []int32{-1, -3, 4, 2},
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angryProfessor(tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("angryProfessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
