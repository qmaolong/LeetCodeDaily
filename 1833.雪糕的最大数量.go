package main

import "sort"

/*
 * @lc app=leetcode.cn id=1833 lang=golang
 *
 * [1833] 雪糕的最大数量
 */

// 63/63 cases passed (228 ms)
// Your runtime beats 59.46 % of golang submissions
// Your memory usage beats 88.59 % of golang submissions (9.2 MB)

// @lc code=start
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for i, v := range costs {
		if coins < v {
			return i
		}
		coins -= v
	}
	if coins > 0 {
		return len(costs)
	}
	return 0
}

// @lc code=end
