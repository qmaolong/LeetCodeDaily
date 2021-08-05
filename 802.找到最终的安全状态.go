package main

import "sort"

/*
 * @lc app=leetcode.cn id=802 lang=golang
 *
 * [802] 找到最终的安全状态
 */

// 112/112 cases passed (124 ms)
// Your runtime beats 37.5 % of golang submissions
// Your memory usage beats 97.92 % of golang submissions (7.4 MB)

// @lc code=start
func eventualSafeNodes(graph [][]int) []int {
	status := make(map[int]int)
	for i := range graph {
		mark(&graph, i, &status)
	}
	res := make([]int, 0)
	for k, v := range status {
		if v == 1 {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}

func mark(graph *[][]int, i int, status *map[int]int) int {
	if _, ok := (*status)[i]; !ok {
		(*status)[i] = 0
	} else {
		return (*status)[i]
	}
	if len((*graph)[i]) == 0 {
		(*status)[i] = 1
		return (*status)[i]
	}
	for _, v := range (*graph)[i] {
		if v == i {
			(*status)[i] = -1
			continue
		}
		if _, ok := (*status)[v]; ok && (*status)[v] == 0 {
			(*status)[i] = -1
			return (*status)[i]
		}
		s := mark(graph, v, status)
		if s == -1 {
			(*status)[i] = -1
			return (*status)[i]
		}
	}
	if (*status)[i] == 0 {
		(*status)[i] = 1
	}
	return (*status)[i]
}

// @lc code=end
