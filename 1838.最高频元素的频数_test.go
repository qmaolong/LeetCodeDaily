package main

import "testing"

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				nums: []int{1, 2, 4},
				k:    5,
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    5,
			},
			want: 2,
		},
		{
			name: "t3",
			args: args{
				nums: []int{3, 9, 16},
				k:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
