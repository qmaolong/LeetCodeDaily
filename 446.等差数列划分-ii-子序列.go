package main

/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	res := 0
	// dp := make([]map[int]int, 0)
	// for i, v := range nums {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		v2 := nums[j]
	// 		dis := v2 - v

	// 	}
	// }
	return res
}

// @lc code=end
