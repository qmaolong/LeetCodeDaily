package main

import (
	"reflect"
	"testing"
)

func Test_distanceK(t *testing.T) {
	type args struct {
		root   *TreeNode
		target *TreeNode
		k      int
	}
	t5 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "t1",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: t5,
					Right: &TreeNode{
						Val: 1,
					},
				},
				target: t5,
				k:      2,
			},
			wantRes: []int{1, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := distanceK(tt.args.root, tt.args.target, tt.args.k); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("distanceK() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
