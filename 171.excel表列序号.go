package main

import "math"

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// 1002/1002 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)

// @lc code=start
func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res += int(columnTitle[i]-'A'+1) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
	}
	return res
}

// @lc code=end
