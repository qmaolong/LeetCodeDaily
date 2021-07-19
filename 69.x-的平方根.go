package main

import "math"

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

//  1017/1017 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 62.72 % of golang submissions (2.2 MB)

// @lc code=start
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// @lc code=end
