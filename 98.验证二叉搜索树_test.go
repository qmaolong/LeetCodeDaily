package main

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "t3",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "t4",
			args: args{
				root: &TreeNode{
					Val: 120,
					Left: &TreeNode{
						Val: 70,
						Left: &TreeNode{
							Val: 50,
							Left: &TreeNode{
								Val: 20,
							},
							Right: &TreeNode{
								Val: 55,
							},
						},
						Right: &TreeNode{
							Val: 100,
							Left: &TreeNode{
								Val: 75,
							},
							Right: &TreeNode{
								Val: 110,
							},
						},
					},
					Right: &TreeNode{
						Val: 140,
						Left: &TreeNode{
							Val: 130,
							Left: &TreeNode{
								Val: 119, //不符合条件
							},
							Right: &TreeNode{
								Val: 135,
							},
						},
						Right: &TreeNode{
							Val: 160,
							Left: &TreeNode{
								Val: 150,
							},
							Right: &TreeNode{
								Val: 200,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
