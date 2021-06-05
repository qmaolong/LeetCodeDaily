package main

/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

// @lc code=start
func reorganizeString(s string) string {
	countMap := make(map[rune]int)
	for _, char := range s {
		countMap[char] += 1
		println(string(char))
	}
	return s
}

// @lc code=end
