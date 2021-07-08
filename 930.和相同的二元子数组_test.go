package main

import "testing"

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{
			name: "t1",
			args: args{
				nums: []int{1, 0, 1, 0, 1},
				//----------1, 1, 2, 2, 3
				goal: 2,
			},
			wantTotal: 4,
		},
		{
			name: "t2",
			args: args{
				nums: []int{0, 0, 0, 0, 0},
				goal: 0,
			},
			wantTotal: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := numSubarraysWithSum(tt.args.nums, tt.args.goal); gotTotal != tt.wantTotal {
				t.Errorf("numSubarraysWithSum() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
