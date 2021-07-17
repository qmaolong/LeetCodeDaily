package main

/*
 * @lc app=leetcode.cn id=1443 lang=golang
 *
 * [1443] 收集树上所有苹果的最少时间
 */

// @lc code=start
func minTime(n int, edges [][]int, hasApple []bool) int {
	suffixMap := make(map[int]int)
	for _, v := range edges {
		suffixMap[v[1]] = v[0]
	}
	routes := make([][]int, 0)
	maxLength := 0
	for i, v := range hasApple {
		if v {
			route := make([]int, 0)
			cursor := i
			for cursor != 0 {
				route = append(route, cursor)
				cursor = suffixMap[cursor]
			}
			routes = append(routes, route)
			if len(route) > maxLength {
				maxLength = len(route)
			}
		}
	}
	total := 0
	for i := 0; i < maxLength; i++ {
		tmpMap := make(map[int]bool)
		for _, route := range routes {
			k := len(route) - i - 1
			if k >= len(route) || k < 0 {
				continue
			}

			if _, ok := tmpMap[route[k]]; ok {
				continue
			}
			total++
			tmpMap[route[k]] = true
		}
	}
	return total * 2
}

// @lc code=end
