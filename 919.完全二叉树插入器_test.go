package main

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	c := Constructor(&root)
	println(c.queue[c.pos].Val)
	v3 := c.Insert(3)
	println(v3)
	println(c.Get_root())
}
