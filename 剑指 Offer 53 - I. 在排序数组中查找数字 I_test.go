package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "t1",
		// 	args: args{
		// 		nums:   []int{5, 7, 7, 8, 8, 10},
		// 		target: 8,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "t2",
		// 	args: args{
		// 		nums:   []int{5, 7, 7, 8, 8, 10},
		// 		target: 6,
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "t3",
		// 	args: args{
		// 		nums:   []int{1},
		// 		target: 0,
		// 	},
		// 	want: 0,
		// },
		{
			name: "t4",
			args: args{
				nums:   []int{1, 1, 2},
				target: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
