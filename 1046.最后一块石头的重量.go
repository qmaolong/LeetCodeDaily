package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

//贪心算法，局部最优解
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 48.81 % of golang submissions (2 MB)

// @lc code=start
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for l := len(stones); l > 1; l = len(stones) {
		a := stones[l-1]
		b := stones[l-2]
		stones = stones[0 : l-2]
		stones = append(stones, a-b)
		sort.Ints(stones)
	}
	return stones[0]
}

// @lc code=end
