package main

/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// 44/44 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 80 % of golang submissions (1.9 MB)

// @lc code=start
func pathInZigZagTree(label int) []int {
	level := 0
	for {
		if label <= (2<<level)-1 {
			break
		}
		level++
	}
	res := make([]int, level+1)
	res[0] = 1
	if label == 1 {
		return res
	}
	var dis int
	if level%2 != 0 {
		dis = 2<<level - 1 - label
	} else {
		dis = label - 2<<(level-1)
	}
	for level != 0 {
		if level%2 != 0 {
			res[level] = (2 << level) - 1 - dis
		} else {
			res[level] = 2<<(level-1) + dis
		}
		dis /= 2
		level--
	}
	return res
}

// @lc code=end
