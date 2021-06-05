package main

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "t1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
