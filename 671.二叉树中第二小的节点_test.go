package main

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
	// r1 := TreeNode{
	// 	Val: 2,
	// 	Left: &TreeNode{
	// 		Val: 5,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 5,
	// 		Left: &TreeNode{
	// 			Val: 5,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// }

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "t1",
		// 	args: args{
		// 		&r1,
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "t2",
		// 	args: args{
		// 		&TreeNode{
		// 			Val: 5,
		// 			Left: &TreeNode{
		// 				Val: 8,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 5,
		// 			},
		// 		},
		// 	},
		// 	want: 8,
		// },
		{
			name: "t3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val: 3,
								Left: &TreeNode{
									Val: 3,
								},
								Right: &TreeNode{
									Val: 3,
								},
							},
							Right: &TreeNode{
								Val: 1,
								Left: &TreeNode{
									Val: 1,
								},
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
						Right: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val: 1,
								Left: &TreeNode{
									Val: 2,
								},
								Right: &TreeNode{
									Val: 1,
								},
							},
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
