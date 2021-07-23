package main

/*
 * @lc app=leetcode.cn id=1893 lang=golang
 *
 * [1893] 检查是否区域内所有整数都被覆盖
 */

//  Your runtime beats 100 % of golang submissions
// Your memory usage beats 20.18 % of golang submissions (2.5 MB)

// @lc code=start
func isCovered(ranges [][]int, left int, right int) bool {
	if len(ranges) == 0 {
		return false
	}
	v := ranges[0]
	if v[0] > right || v[1] < left {
		return isCovered(ranges[1:], left, right)
	} else if v[1] >= right && v[0] <= left {
		return true
	}
	if v[0] > left {
		cover := isCovered(ranges[1:], left, v[0]-1)
		if !cover {
			return false
		}
	}
	if v[1] < right {
		cover := isCovered(ranges[1:], v[1]+1, right)
		if !cover {
			return false
		}
	}
	return true
}

// @lc code=end
