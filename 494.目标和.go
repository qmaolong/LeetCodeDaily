package main

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	return findTargetSumWays(nums[1:], target-nums[0]) + findTargetSumWays(nums[1:], target+nums[0])
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
