package main

/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// 52/52 cases passed (80 ms)
// Your runtime beats 16.32 % of golang submissions
// Your memory usage beats 13.68 % of golang submissions (7.6 MB)

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	queue := make([][]int, 0)
	m := make(map[int][][]int)
	for _, v := range times {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = make([][]int, 0)
		}
		m[v[0]] = append(m[v[0]], v)
	}
	queue = append(queue, []int{k, 0})
	sead := make(map[int]int)
	sead[k] = 0
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		nexts := m[q[0]]
		for _, v := range nexts {
			time := q[1] + v[2]
			if _, ok := sead[v[1]]; !ok || time < sead[v[1]] {
				queue = append(queue, []int{v[1], time})
				sead[v[1]] = time
			}
		}
	}
	if len(sead) != n {
		return -1
	}
	max := 0
	for _, v := range sead {
		if v > max {
			max = v
		}
	}
	return max
}

// @lc code=end
