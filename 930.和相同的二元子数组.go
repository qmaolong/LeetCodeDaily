package main

/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 */

//  60/60 cases passed (28 ms)
// Your runtime beats 98 % of golang submissions
// Your memory usage beats 52 % of golang submissions (8.1 MB)

// @lc code=start
func numSubarraysWithSum(nums []int, goal int) (total int) {
	prefixSum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			prefixSum[i] = v
		} else {
			prefixSum[i] = v + prefixSum[i-1]
		}
	}
	m := make(map[int]int)
	for _, v := range prefixSum {
		if v == goal {
			total++
		}
		total += m[v-goal]
		m[v]++
	}
	return
}

// @lc code=end
