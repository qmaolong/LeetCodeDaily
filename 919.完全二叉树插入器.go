package main

/*
 * @lc app=leetcode.cn id=919 lang=golang
 *
 * [919] 完全二叉树插入器
 */

//  83/83 cases passed (8 ms)
// Your runtime beats 95 % of golang submissions
// Your memory usage beats 25 % of golang submissions (6.7 MB)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
	pos   int
}

func Constructor(root *TreeNode) CBTInserter {
	cbt := CBTInserter{root: root, pos: -1}
	//BFS
	q := make([]*TreeNode, 0)
	q = append(q, root)
	c := 0
	for c < len(q) {
		n := q[c]
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
		if cbt.pos == -1 && (n.Left == nil || n.Right == nil) {
			cbt.pos = c
		}
		c++
	}
	cbt.queue = q
	return cbt
}

func (this *CBTInserter) Insert(v int) int {
	point := TreeNode{Val: v}
	parentNode := this.GetPositionNode()
	if parentNode.Left == nil {
		parentNode.Left = &point
	} else {
		parentNode.Right = &point
		this.pos += 1
	}
	this.queue = append(this.queue, &point)
	return parentNode.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func (this *CBTInserter) GetPositionNode() *TreeNode {
	return this.queue[this.pos]
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
// @lc code=end
