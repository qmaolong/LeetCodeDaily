package main

import "sort"

/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

//O(n^2)
//  241/241 cases passed (28 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 71.95 % of golang submissions (3.2 MB)

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 2; i < len(nums); i++ {
		left := 0
		right := i - 1
		for left != right {
			if nums[left]+nums[right] > nums[i] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// @lc code=end
