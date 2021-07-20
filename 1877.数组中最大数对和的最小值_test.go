package main

import "testing"

func Test_minPairSum(t *testing.T) {
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
				nums: []int{3, 5, 2, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPairSum(tt.args.nums); got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
