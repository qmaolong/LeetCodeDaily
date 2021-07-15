package main

/*
 * @lc app=leetcode.cn id=1860 lang=golang
 *
 * [1860] 增长的内存泄露
 */

//  82/82 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 60.76 % of golang submissions (1.9 MB)

// @lc code=start
func memLeak(memory1 int, memory2 int) []int {
	crashTime := 1
	for {
		if memory1 < crashTime && memory2 < crashTime {
			break
		}
		if memory2 > memory1 {
			memory2 -= crashTime
		} else {
			memory1 -= crashTime
		}
		crashTime++
	}
	return []int{crashTime, memory1, memory2}
}

// @lc code=end
