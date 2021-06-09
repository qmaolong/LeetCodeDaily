package main

/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// Your runtime beats 61.11 % of golang submissions
// Your memory usage beats 5.55 % of golang submissions (24.8 MB)

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	MOD := 1000000007
	dp := make([][][]int, len(profit))
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}

	for i := range dp {
		g := group[i]
		p := profit[i]
		for j := range dp[i] {
			for k := range dp[i][j] {
				if i == 0 {
					if j >= g && k <= p {
						dp[i][j][k] = 1
					}
				} else {
					if j >= g && k <= p {
						g1 := j - g
						dp[i][j][k] = dp[i-1][g1][0] + 1 + dp[i-1][j][k]
					} else if j >= g {
						p1 := k - p
						g1 := j - g
						dp[i][j][k] = dp[i-1][g1][p1] + dp[i-1][j][k]
					} else {
						dp[i][j][k] = dp[i-1][j][k]
					}
					dp[i][j][k] = dp[i][j][k] % MOD
				}
			}
		}
	}
	res := dp[len(profit)-1][n][minProfit]
	if minProfit == 0 {
		res++
	}
	return res % MOD
}

// @lc code=end

//x    1    2    3    4    5    6    7    8    9    10
//2,6  0,1  1,0  1,1  1,2  1,3  1,4  1,5  1,6  1,7	1,8
//3,7  0,1  1,0  2,x  2,x  3,x  3,x  3,x  3,x  3,x  3,x
//5,8  0,1  1,0  2,x  2,x  4,x 3+1=0 3+1+1

//y=dp[i-1][j-g]

//x    1   2   3   4   5
//2,2  0,0 1,0 1,0 1,0 1,0
//2,3  0,0 2,1 2,1 3,2 3,2
