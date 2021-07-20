package main

import "sort"

/*
 * @lc app=leetcode.cn id=1877 lang=golang
 *
 * [1877] 数组中最大数对和的最小值
 */

//  37/37 cases passed (292 ms)
// Your runtime beats 92.5 % of golang submissions
// Your memory usage beats 51.67 % of golang submissions (9.1 MB)

// @lc code=start
func minPairSum(nums []int) int {
	res := 0
	i := 0
	l := len(nums)
	sort.Ints(nums)
	for {
		sum := nums[i] + nums[l-i-1]
		if sum > res {
			res = sum
		}
		if i+1 == l-i-1 {
			break
		}
		i++
	}
	return res
}

// @lc code=end
