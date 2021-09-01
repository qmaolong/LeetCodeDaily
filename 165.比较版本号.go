package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

//  81/81 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 79.15 % of golang submissions (1.9 MB)

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	l := len(arr1)
	if len(arr2) > l {
		l = len(arr2)
	}
	for i := 0; i < l; i++ {
		n1 := 0
		n2 := 0
		if len(arr1)-1 >= i {
			n1, _ = strconv.Atoi(arr1[i])
		}
		if len(arr2)-1 >= i {
			n2, _ = strconv.Atoi(arr2[i])
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
	}
	return 0
}

// @lc code=end
