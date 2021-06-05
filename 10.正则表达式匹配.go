package main

import "strings"

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	sCursor := 0
	pCursor := 0

	matchChar := func(a byte, b string) bool {
		return string(a) == b || b == "."
	}

	pSlice := make([]string, 0)
	for i, v := range p {
		if v == '*' {
			pSlice = append(pSlice, string(p[i-1])+string(v))
		} else {
			pSlice = append(pSlice, string(v))
		}
	}
	for pCursor < len(pSlice) {
		pItem := pSlice[pCursor]
		if strings.HasSuffix(pItem, "*") {
			matchCount := 0
			for matchChar(s[sCursor+matchCount], pItem) {
				matchCount++
			}
			for i := 0; i < matchCount; i++ {
				//if isMatch() {
				//
				//}
			}
		} else if sCursor < len(s) && matchChar(s[sCursor], pItem) {
			sCursor++
			pCursor++
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

//@date 20210429失败记录
//@reason 开始的思路错误导致后面的努力白费
//@summary
//1.将正则拆分成数组
//2.记录每个正则元素匹配的字符传位置
//3.计算每个正则匹配的字符可否覆盖整个输入的字符串
//
// func isMatch(s string, p string) bool {
// 	sCursor := 0
// 	pCursor := 0

// 	matchChar := func(a byte, b byte) bool {
// 		return a == b || b == byte('.')
// 	}
// 	for pCursor < len(p) && sCursor < len(s) {
// 		c := p[pCursor]
// 		if pCursor+1 < len(p) && p[pCursor+1] == byte('*') {
// 			for sCursor < len(s) && matchChar(s[sCursor], c) {
// 				sCursor++
// 			}
// 			pCursor++
// 		} else if matchChar(s[sCursor], c) {
// 			sCursor++
// 		} else {
// 			return false
// 		}
// 		pCursor++
// 	}
// 	return sCursor == len(s) && pCursor == len(p)
// }
