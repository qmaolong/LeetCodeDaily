package main

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

//  Your runtime beats 16.28 % of golang submissions
// Your memory usage beats 96.12 % of golang submissions (4.6 MB)

// @lc code=start
func frequencySort(s string) string {
	freq := make(map[byte]int)
	for _, v := range s {
		freq[byte(v)] = freq[byte(v)] + 1
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return strings.IndexByte(s, b[i]) > strings.IndexByte(s, b[j])
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}

// @lc code=end
