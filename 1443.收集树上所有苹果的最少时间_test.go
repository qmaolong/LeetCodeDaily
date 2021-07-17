package main

import "testing"

func Test_minTime(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		hasApple []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				n:        7,
				edges:    [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				hasApple: []bool{false, false, true, false, true, true, false},
			},
			want: 8,
		},
		{
			name: "t2",
			args: args{
				n:        7,
				edges:    [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				hasApple: []bool{false, false, true, false, false, true, false},
			},
			want: 6,
		},
		{
			name: "t3",
			args: args{
				n:        7,
				edges:    [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				hasApple: []bool{false, false, false, false, false, false, false},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTime(tt.args.n, tt.args.edges, tt.args.hasApple); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
