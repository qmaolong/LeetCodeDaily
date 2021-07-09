package main

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "t1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantRes: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := majorityElement(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("majorityElement() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
