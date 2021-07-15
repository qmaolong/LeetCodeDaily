package main

import "sort"

/*
 * @lc app=leetcode.cn id=1846 lang=golang
 *
 * [1846] 减小和重新排列数组后的最大元素
 */

// 49/49 cases passed (80 ms)
// Your runtime beats 76.47 % of golang submissions
// Your memory usage beats 53.78 % of golang submissions (9.5 MB)

// @lc code=start
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		sub := v - arr[i-1]
		if sub > 1 {
			arr[i] = arr[i-1] + 1
		} else if sub < -1 {
			return arr[i-1]
		}
	}
	return arr[len(arr)-1]
}

// @lc code=end
