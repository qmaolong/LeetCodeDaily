package main

/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 * 前缀和+哈希表
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	sums := make([]int, len(nums))
	targets := make(map[int]int)
	for i, v := range nums {
		sum := v
		if i > 0 {
			sum += sums[i-1]
		}
		reminder := sum % k
		if i > 0 && reminder == 0 {
			return true
		}
		sums[i] = sum
		if targets[reminder] != 0 && i-targets[reminder] > 0 {
			return true
		}
		if targets[reminder] == 0 {
			targets[reminder] = i + 1
		}
	}
	return false
}

// @lc code=end
