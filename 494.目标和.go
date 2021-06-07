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
