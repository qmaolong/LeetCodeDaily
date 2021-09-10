package main

/*
 * @lc app=leetcode.cn id=1894 lang=golang
 *
 * [1894] 找到需要补充粉笔的学生编号
 */

// 72/72 cases passed (120 ms)
// Your runtime beats 94.23 % of golang submissions
// Your memory usage beats 69.23 % of golang submissions (8.1 MB)

// @lc code=start
func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k %= sum
	for i, v := range chalk {
		if k < v {
			return i
		}
		k -= v
	}
	return 0
}

// @lc code=end
