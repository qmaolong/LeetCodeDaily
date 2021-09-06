package main

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

//  46/46 cases passed (32 ms)
// Your runtime beats 84.4 % of golang submissions
// Your memory usage beats 82.29 % of golang submissions (6.7 MB)

// @lc code=start
func search1(nums []int, target int) int {
	c1 := 0
	c2 := len(nums) - 1
	for c1 <= c2 {
		c := (c1 + c2) / 2
		if nums[c] == target {
			return c
		} else if nums[c] > target {
			c2 = c - 1
		} else {
			c1 = c + 1
		}
	}
	return -1
}

// @lc code=end
