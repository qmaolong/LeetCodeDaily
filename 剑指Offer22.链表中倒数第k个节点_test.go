package main

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	node4 := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
		},
	}
	node1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &node4,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "t1",
			args: args{
				head: &node1,
				k:    2,
			},
			want: &node4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
