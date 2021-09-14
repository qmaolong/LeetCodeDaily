package main

import "sort"

/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// 31/31 cases passed (16 ms)
// Your runtime beats 79.31 % of golang submissions
// Your memory usage beats 70.34 % of golang submissions (7.8 MB)

// @lc code=start
func findLongestWord(s string, dictionary []string) (res string) {
	sort.Strings(dictionary)
	for _, v := range dictionary {
		if len(v) <= len(res) {
			continue
		}
		c := 0
		for i := 0; i < len(s); i++ {
			if s[i] == v[c] {
				c++
			}
			if c == len(v) {
				res = v
				break
			}
		}
	}
	return
}

// @lc code=end
