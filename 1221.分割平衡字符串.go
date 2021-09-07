package main

/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

//  40/40 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)

// @lc code=start
func balancedStringSplit(s string) int {
	r := 0
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'L' {
			n--
		} else {
			n++
		}
		if n == 0 {
			r++
		}
	}
	return r
}

// @lc code=end
