package main

/*
 * @lc app=leetcode.cn id=1227 lang=golang
 *
 * [1227] 飞机座位分配概率
 */

// 100/100 cases passed (8 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)

// @lc code=start
func nthPersonGetsNthSeat(n int) float64 {
	prev := float64(1)
	for i := float64(2); i <= float64(n); i++ {
		prev = 1/i + (i-2)/i*prev
	}
	return prev
}

// @lc code=end
