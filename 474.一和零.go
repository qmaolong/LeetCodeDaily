package main

import "strings"

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs))
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := range dp {
		s := strs[i]
		oneCount := strings.Count(s, "1")
		zeroCount := len(s) - oneCount
		for j := range dp[i] {
			for k := range dp[i][j] {
				if zeroCount > j || oneCount > k {
					if i == 0 {
						dp[i][j][k] = 0
					} else {
						dp[i][j][k] = dp[i-1][j][k]
					}
				} else {
					if i == 0 {
						dp[i][j][k] = 1
					} else {
						leftZero := j - zeroCount
						leftOne := k - oneCount
						dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][leftZero][leftOne]+1)
					}
				}
			}
		}
	}
	return dp[len(strs)-1][m][n]
}

// @lc code=end
