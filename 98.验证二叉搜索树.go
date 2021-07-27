package main

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

//  Your runtime beats 95.6 % of golang submissions
// Your memory usage beats 98.41 % of golang submissions (5.1 MB)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min *int, max *int) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if min != nil && root.Left.Val <= *min {
			return false
		}
		if max != nil && root.Left.Val >= *max {
			return false
		}
		if !isValid(root.Left, min, &root.Val) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if min != nil && root.Right.Val <= *min {
			return false
		}
		if max != nil && root.Right.Val >= *max {
			return false
		}
		if !isValid(root.Right, &root.Val, max) {
			return false
		}
	}
	return true
}

// @lc code=end
