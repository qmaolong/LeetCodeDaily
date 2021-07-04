package main

/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

//  49/49 cases passed (28 ms)
// Your runtime beats 91.28 % of golang submissions
// Your memory usage beats 53.69 % of golang submissions (6.5 MB)

// @lc code=start
func findErrorNums(nums []int) []int {
	sorted := make([]int, len(nums)+1)
	res := make([]int, 2)
	for _, v := range nums {
		if sorted[v] == v {
			res[0] = v
		}
		sorted[v] = v
	}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] != i {
			res[1] = i
			break
		}
	}
	return res
}

// @lc code=end
