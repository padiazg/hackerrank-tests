package main

import "testing"

func Test_designerPdfViewer(t *testing.T) {
	type args struct {
		h    []int32
		word string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "input",
			args: args{
				h:    []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5},
				word: "torn",
			},
			want: 8,
		},
		{
			name: "input00",
			args: args{
				h:    []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
				word: "abc",
			},
			want: 9,
		},
		{
			name: "input00",
			args: args{
				h:    []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
				word: "zaba",
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := designerPdfViewer(tt.args.h, tt.args.word); got != tt.want {
				t.Errorf("designerPdfViewer() = %v, want %v", got, tt.want)
			}
		})
	}
}
