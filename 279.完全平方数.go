package main

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

//  Your runtime beats 43.59 % of golang submissions
//  Your memory usage beats 75.72 % of golang submissions (6.1 MB)
// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	x := 1
	for i := 1; x <= n; i++ {
		if n == x {
			return 1
		}
		dp[x] = 1
		for j := x + 1; j <= n; j++ {
			s := dp[j-x] + 1
			if dp[j] == 0 || s < dp[j] {
				dp[j] = s
			}
		}
		x = (i + 1) * (i + 1)
	}
	return dp[n]
}

// @lc code=end

//x 1   2   3   4   5
//1 1   2   3   4   5
//4 1   2   3   1   2
//
