package main

/*
 * @lc app=leetcode.cn id=1711 lang=golang
 *
 * [1711] 大餐计数
 */

// 70/70 cases passed (184 ms)
// Your runtime beats 53.45 % of golang submissions
// Your memory usage beats 39.66 % of golang submissions (9.6 MB)

// @lc code=start
func countPairs(deliciousness []int) (ans int) {
	max := deliciousness[0]
	for i := 1; i < len(deliciousness); i++ {
		if deliciousness[i] > max {
			max = deliciousness[i]
		}
	}
	m := make(map[int]int)
	max <<= 1
	for _, v := range deliciousness {
		for i := 1; i <= max; i <<= 1 {
			ans += m[i-v]
		}
		m[v]++
	}
	return ans % 1000000007
}

// @lc code=end
