package main

/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// Your runtime beats 100 % of golang submissions
// Your memory usage beats 26.83 % of golang submissions (3.2 MB)

// @lc code=start
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	harf := sum / 2

	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, harf+1)
	}

	max := 0
	for i := range dp {
		stone := stones[i]
		for j := range dp[i] {
			if i == 0 {
				if j == stone {
					dp[i][j] = 1
				}
			} else {
				if dp[i-1][j] == 1 || stone == j {
					dp[i][j] = 1
					max = j
					continue
				}
				t := j - stone
				if t < 0 {
					continue
				}
				if dp[i-1][t] == 1 {
					dp[i][j] = 1
				}
			}
			if dp[i][j] == 1 {
				max = j
			}
		}
	}
	return sum - (max * 2)
}

// @lc code=end
