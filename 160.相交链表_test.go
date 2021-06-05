package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	v4 := ListNode{
		Val:  4,
		Next: nil,
	}
	v2 := ListNode{
		Val:  2,
		Next: &v4,
	}
	v1 := ListNode{
		Val:  1,
		Next: &v2,
	}
	v9 := ListNode{
		Val:  9,
		Next: &v1,
	}
	v11 := ListNode{
		Val:  1,
		Next: &v9,
	}
	v3 := ListNode{
		Val:  3,
		Next: &v2,
	}
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "t1",
			args: args{
				headA: &v11,
				headB: &v3,
			},
			want: &v2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
