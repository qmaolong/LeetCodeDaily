package main

/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	m := make(map[int]map[int]int, n)
	for _, v := range flights {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = make(map[int]int, n)
		}
		m[v[0]][v[1]] = v[2]
	}
	type Item struct {
		node  int
		deep  int
		price int
	}
	res := -1
	q := make([]Item, 0)
	q = append(q, Item{src, 0, 0})
	for len(q) != 0 && q[0].deep <= k {
		it := q[0]
		q = q[1:]
		nexts := m[it.node]
		for k2, v := range nexts {
			if k2 == dst {
				price := it.price + v
				if res < 0 || price < res {
					res = price
				}
			} else {
				q = append(q, Item{k2, it.deep + 1, it.price + v})
			}
		}
	}
	return res
}

// @lc code=end
