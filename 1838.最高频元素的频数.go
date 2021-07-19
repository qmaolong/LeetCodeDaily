package main

import "sort"

/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] 最高频元素的频数
 */

//  71/71 cases passed (204 ms)
// Your runtime beats 96.3 % of golang submissions
// Your memory usage beats 30.86 % of golang submissions (9.2 MB)

// @lc code=start
func maxFrequency(nums []int, k int) int {
	l := len(nums)
	if l == 1 {
		return 1
	}
	sort.Ints(nums)
	prefixSum := make([]int, l)
	prefixSum[0] = nums[0]
	for i := 1; i < l; i++ {
		prefixSum[i] = nums[i] + prefixSum[i-1]
	}
	max := 0
	cur1 := 0
	cur2 := 1
	for {
		s2 := prefixSum[cur2]
		s1 := 0
		if cur1 > 0 {
			s1 = prefixSum[cur1-1]
		}
		step := cur2 - cur1 + 1
		if s2-s1+k < nums[cur2]*step {
			cur1++
		} else {
			if step > max {
				max = step
			}
			cur2++
		}
		if cur2 >= l {
			break
		}
	}
	return max

}

// @lc code=end
