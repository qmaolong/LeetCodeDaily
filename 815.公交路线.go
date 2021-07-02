package main

import "fmt"

/*
 * @lc app=leetcode.cn id=815 lang=golang
 *
 * [815] 公交路线
 */

//44/45 cases passed (N/A)####大数据量超时

// @lc code=start
func numBusesToDestination(routes [][]int, source int, target int) int {
	type option struct {
		point int
		step  int
	}
	queue := make([]option, 0)
	queue = append(queue, option{
		source,
		0,
	})
	routesMap := make([]map[int]bool, len(routes))
	for i, r := range routes {
		m := make(map[int]bool)
		for _, p := range r {
			m[p] = true
		}
		routesMap[i] = m
	}
	sead := make(map[int]bool)
	sead[source] = true
	for len(queue) > 0 {
		fmt.Println(len(queue))
		op := queue[0]
		queue = queue[1:]
		if op.point == target {
			return op.step
		}
		for i, r := range routesMap {
			if r[op.point] {
				if r[target] {
					return op.step + 1
				}
				for _, p := range routes[i] {
					if !sead[p] && p != op.point {
						sead[p] = true
						queue = append(queue, option{
							p,
							op.step + 1,
						})
					}
				}
			}
		}
	}
	return -1
}

// @lc code=end
