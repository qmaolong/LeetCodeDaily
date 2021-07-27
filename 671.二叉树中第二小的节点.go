package main

/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
 */

// Your runtime beats 100 % of golang submissions
// Your memory usage beats 98.9 % of golang submissions (1.9 MB)

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
	if leftFirstBiggerValue != -1 && (res == -1 || leftFirstBiggerValue < res) {
		res = leftFirstBiggerValue
	}
	rightFirstBiggerValue := findFirstBiggerValue(root.Right, root.Val)
	if rightFirstBiggerValue != -1 && (res == -1 || rightFirstBiggerValue < res) {
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
		if left != -1 && (res == -1 || left < res) {
			res = left
		}
		righ := findFirstBiggerValue(node.Right, target)
		if righ != -1 && (res == -1 || righ < res) {
			res = righ
		}
		return res
	} else {
		return -1
	}
}

// @lc code=end

//1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1
//            1
//        1          3
//   1        1    3    4
// 3    1,   1 1, 3 8, 4 8
//3 3, 1 6, 2 1
