package main

/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	total := 0
	l := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, v := range nums {
			c += v >> i & 1
		}
		total += c * (l - c)
	}
	return total
}

// @lc code=end
