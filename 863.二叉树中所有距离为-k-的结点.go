package main

/*
 * @lc app=leetcode.cn id=863 lang=golang
 *
 * [863] 二叉树中所有距离为 K 的结点
 */

// 57/57 cases passed (4 ms)
// Your runtime beats 45.83 % of golang submissions
// Your memory usage beats 12.5 % of golang submissions (3.5 MB)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) (res []int) {
	m := make(map[*TreeNode]*TreeNode)
	dfs(root, m)
	history := make(map[*TreeNode]bool)
	distance(target, k, 0, m, history, &res)
	return
}

func distance(root *TreeNode, k int, deep int, m map[*TreeNode]*TreeNode, history map[*TreeNode]bool, res *[]int) {
	if _, has := history[root]; has {
		return
	}
	history[root] = true
	if deep == k {
		*res = append(*res, root.Val)
	}
	if root.Left != nil {
		distance(root.Left, k, deep+1, m, history, res)
	}
	if root.Right != nil {
		distance(root.Right, k, deep+1, m, history, res)
	}
	if _, ok := m[root]; ok {
		distance(m[root], k, deep+1, m, history, res)
	}
}

func dfs(node *TreeNode, m map[*TreeNode]*TreeNode) {
	if node.Left != nil {
		m[node.Left] = node
		dfs(node.Left, m)
	}
	if node.Right != nil {
		m[node.Right] = node
		dfs(node.Right, m)
	}
}

// @lc code=end
