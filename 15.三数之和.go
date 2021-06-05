package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	c1plus := func(c1 *int) {
		*c1++
		for (*c1) < len(nums) && nums[*c1] == nums[*c1-1] {
			*c1++
		}
	}

	c2minus := func(c2 *int) {
		*c2--
		for (*c2) > 0 && nums[*c2] == nums[*c2+1] {
			*c2--
		}
	}
	c0 := 0
	length := len(nums)
	for c0 < length {
		c1 := c0 + 1
		c2 := length - 1
		num1 := nums[c0]
		for c1 < c2 {
			if nums[c1]+nums[c2]+num1 < 0 {
				c1plus(&c1)
			} else if nums[c1]+nums[c2]+num1 > 0 {
				c2minus(&c2)
			} else {
				result = append(result, []int{num1, nums[c1], nums[c2]})
				c1plus(&c1)
				c2minus(&c2)
			}
		}
		c0++
		for c0 < length && nums[c0] == nums[c0-1] {
			c0++
		}
	}
	return result
}

// @lc code=end
