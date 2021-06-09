package main

import "testing"

func Test_match(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
		x      int
		y      int
	}
	matrix := [][]int{
		{1, 4, 7, 11, 15, 16},
		{2, 5, 8, 12, 19, 20},
		{3, 6, 9, 16, 22, 25},
		{10, 13, 14, 17, 24, 27},
		{18, 21, 23, 26, 30, 35},
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				matrix: matrix,
				target: 16,
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				matrix: matrix,
				target: 100,
			},
			want: false,
		},
		{
			name: "t3",
			args: args{
				matrix: matrix,
				target: 35,
			},
			want: true,
		},
		{
			name: "t4",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7, 9},
					{2, 4, 6, 8, 10},
					{11, 13, 15, 17, 19},
					{12, 14, 16, 18, 20},
					{21, 22, 23, 24, 25},
				},
				target: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.matrix, tt.args.target, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
