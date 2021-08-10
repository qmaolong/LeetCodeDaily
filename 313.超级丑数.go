package main

/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	for i, v := range primes {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[j] = v
				} else {
					dp[j] = v * dp[j-1]
				}
			} else {
				if j == 0 {
					dp[j] = min(v, dp[j])
				} else {
					dp[j] = v * min(dp[j], dp[j-1])
				}
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
