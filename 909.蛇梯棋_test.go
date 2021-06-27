package main

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "t1",
		// 	args: args{
		// 		board: [][]int{
		// 			{-1, -1, -1, -1, -1, -1},
		// 			{-1, -1, -1, -1, -1, -1},
		// 			{-1, -1, -1, -1, -1, -1},
		// 			{-1, 35, -1, -1, 13, -1},
		// 			{-1, -1, -1, -1, -1, -1},
		// 			{-1, 15, -1, -1, -1, -1},
		// 		},
		// 	},
		// 	want: 4,
		// },
		{
			name: "t2",
			args: args{
				board: [][]int{
					{-1, -1, -1}, {-1, 9, 8}, {-1, 8, 9},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.args.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
