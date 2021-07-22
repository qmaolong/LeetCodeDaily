package main

import (
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	n5 := Node{
		Val:  1,
		Next: nil,
	}
	n4 := Node{
		Val:  10,
		Next: &n5,
	}
	n3 := Node{
		Val:  11,
		Next: &n4,
	}
	n2 := Node{
		Val:  13,
		Next: &n3,
	}
	n1 := Node{
		Val:  7,
		Next: &n2,
	}
	n1.Random = nil
	n2.Random = &n1
	n3.Random = &n5
	n4.Random = &n3
	n5.Random = &n1

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "t1",
			args: args{
				head: &n1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList(tt.args.head)
			c1 := tt.args.head
			c2 := got
			for c1 != nil || c2 != nil {
				if c1.Val != c2.Val {
					t.Errorf("复制出错")
				}
				if !nodeEqual(c1.Random, c2.Random) {
					t.Error("随机指针出错")
				}
				c1 = c1.Next
				c2 = c2.Next
			}
		})
	}
}

func nodeEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil && node1.Val == node2.Val {
		return true
	}
	return false
}
