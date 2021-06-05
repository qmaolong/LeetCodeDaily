package main

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z != 0 {
		if z%2 == 1 {
			count++
		}
		z = z >> 1
	}
	return count
}

// @lc code=end
