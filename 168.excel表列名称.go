package main

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// 18/18 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)

// @lc code=start
func convertToTitle(columnNumber int) string {
	var arr = []byte{}
	offset := 'A' - 1
	for columnNumber > 0 {
		e := columnNumber % 26
		columnNumber /= 26
		if e == 0 {
			e = 26
			columnNumber -= 1
		}
		arr = append(arr, byte(e+int(offset)))
	}
	var reverse = []byte{}
	l := len(arr)
	for i := range arr {
		reverse = append(reverse, arr[l-i-1])
	}
	return string(reverse)
}

// @lc code=end
