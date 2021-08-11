package main

import "testing"

func Test_numberOfArithmeticSlices1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices1(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
