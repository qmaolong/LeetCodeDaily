package main

/*
 * @lc app=leetcode.cn id=1713 lang=golang
 *
 * [1713] 得到子序列的最少操作次数
 */

//  Your runtime beats 31.33 % of golang submissions
//  Your memory usage beats 71.11 % of golang submissions (12.3 MB)

// @lc code=start
func minOperations(target []int, arr []int) int {
	m := make(map[int]int)
	for i, v := range target {
		m[v] = i
	}
	d := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		item := arr[i]
		if _, has := m[item]; !has {
			continue
		}
		l := len(d)
		v := m[item]
		if l == 0 || d[l-1] < v {
			d = append(d, v)
			continue
		}
		c1 := 0
		c2 := l - 1
		var c int
		for c1 <= c2 {
			c = (c1 + c2) / 2
			if v < d[c] && (c == 0 || v > d[c-1]) {
				d[c] = v
			}
			if d[c] > v {
				c2 = c - 1
			} else if d[c] < v {
				c1 = c + 1
			} else {
				break
			}
		}
	}
	return len(target) - len(d)
}

// @lc code=end
