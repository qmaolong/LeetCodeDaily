package main

/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}
	res := -1
	leftFirstBiggerValue := findFirstBiggerValue(root.Left, root.Val)
	if leftFirstBiggerValue > res {
		res = leftFirstBiggerValue
	}
	rightFirstBiggerValue := findFirstBiggerValue(root.Right, root.Val)
	if rightFirstBiggerValue > res {
		res = rightFirstBiggerValue
	}
	return res
}

func findFirstBiggerValue(node *TreeNode, target int) int {
	if node.Val > target {
		return node.Val
	} else if node.Left != nil {
		res := -1
		left := findFirstBiggerValue(node.Left, target)
		if left > res {
			res = left
		}
		righ := findFirstBiggerValue(node.Right, target)
		if righ > res {
			res = righ
		}
		return res
	} else {
		return -1
	}
}

// @lc code=end
