package main

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	result := 0
	for x != 0 {
		i := x % 10
		x = x / 10
		result = 10*result + i
	}
	if result > (1<<31)-1 || result < -1<<31 {
		result = 0
	}
	return result
}

// @lc code=end
