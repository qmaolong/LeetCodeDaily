package main

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				points: [][]int{
					{1, 1}, {2, 2}, {3, 3},
				},
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				points: [][]int{
					{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4},
				},
			},
			want: 4,
		},
		{
			name: "t3",
			args: args{
				points: [][]int{
					{0, 0},
				},
			},
			want: 1,
		},
		{
			name: "t4",
			args: args{
				points: [][]int{
					{4, 5}, {4, -1}, {4, 0},
				},
			},
			want: 3,
		},
		{
			name: "t5",
			args: args{
				points: [][]int{
					{0, 0}, {4, 5}, {7, 8}, {8, 9}, {5, 6}, {3, 4}, {1, 1},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
