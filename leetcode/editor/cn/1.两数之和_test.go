package leecode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "t1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			wantRes: []int{0, 1},
		},
		{
			name: "t2",
			args: args{
				nums:   []int{0, 4, 3, 0},
				target: 0,
			},
			wantRes: []int{0, 3},
		},
		{
			name: "t3",
			args: args{
				nums:   []int{1, 3, 4, 2},
				target: 6,
			},
			wantRes: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("twoSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
