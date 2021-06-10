package main

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start

//动态规划实现
// Your runtime beats 5 % of golang submissions
// Your memory usage beats 8.8 % of golang submissions (12.6 MB)
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	} else if len(coins) == 0 {
		return 0
	} else if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return 1
		}
		return 0
	}
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		coin := coins[i]
		for j := range dp[i] {
			if j == 0 {
				continue
			}
			if i == 0 {
				if j%coin == 0 {
					dp[i][j] = 1
				}
			} else {
				total := 0
				for k := 0; k*coin <= j; k++ {
					l := j - k*coin
					if l == 0 {
						total += 1
					} else {
						total += dp[i-1][l]
					}
				}
				dp[i][j] = total
			}
		}
	}
	res := dp[len(coins)-1][amount]
	if amount == 0 {
		res += 1
	}
	return res
}

// @lc code=end

//x  1   2   3   4   5
//1  1   1   1   1   1
//2  1   2   2   3   3
//5  1   2   2   3   3+1
