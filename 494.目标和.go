package main

import "fmt"

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	dp := make([][1001]int, len(nums))

	for i := range dp {
		num := nums[i]
		for j := range dp[i] {
			if i == 0 {
				if j == num || j == -num {
					dp[i][j] = 1
				}
			} else {
				t1 := j - num
				if t1 < 0 {
					t1 = -t1
				}
				t2 := j + num
				if t2 < 0 {
					t2 = -t2
				}
				if t1 == 1001 || t2 == 1001 {
					fmt.Print(t1)
				}
				dp[i][j] = dp[i-1][t1] + dp[i-1][t2]
			}
		}
	}
	return dp[len(nums)-1][target]
}

// @lc code=end

/* 常规解法
时间复杂度：O（2^n）
Your runtime beats 25.91 % of golang submissions
Your memory usage beats 83.15 % of golang submissions (2.2 MB)
-------------------------
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	return findTargetSumWays(nums[1:], target-nums[0]) + findTargetSumWays(nums[1:], target+nums[0])
} */
