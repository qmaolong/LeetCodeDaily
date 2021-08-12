package main

/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

//边界：dp[i][i]=1
//状态转移：
//1.when dp[i]==dp[j] then dp[i][j]=dp[i+1][j]+2=dp[i][j-1]
//2.when dp[i]!=dp[j] then dp[i][j]=Max(dp[i+1][j],dp[i][j-1])
// 86/86 cases passed (24 ms)
// Your runtime beats 88.2 % of golang submissions
// Your memory usage beats 6.23 % of golang submissions (18.9 MB)

// @lc code=start
func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := l - 1; i >= 0; i-- {
		dp[i] = make([]int, l)
		for j := i; j < l; j++ {
			if i == j {
				dp[i][j] = 1
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = dp[i+1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[0][l-1]
}

// @lc code=end
