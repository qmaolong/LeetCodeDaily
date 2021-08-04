package main

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				n:     4,
				k:     2,
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				times: [][]int{{1, 2, 1}},
				n:     2,
				k:     1,
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
				n:     3,
				k:     1,
			},
			want: 3,
		},
		{
			name: "t4",
			args: args{
				times: [][]int{{1, 5, 66}, {3, 5, 55}, {4, 3, 29}, {1, 2, 9}, {3, 4, 10}, {3, 1, 3}, {2, 3, 78}, {1, 4, 98}, {4, 5, 21}, {5, 2, 19}, {5, 1, 76}, {4, 1, 65}, {3, 2, 27}, {5, 3, 23}, {5, 4, 12}, {2, 1, 36}, {4, 2, 75}, {2, 4, 11}, {1, 3, 30}, {2, 5, 8}},
				n:     5,
				k:     1,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
