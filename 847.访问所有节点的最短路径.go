package main

/*
 * @lc app=leetcode.cn id=847 lang=golang
 *
 * [847] 访问所有节点的最短路径
 */
type VNode struct {
	Val       int
	MaxLength int
	Link      []int
}

// @lc code=start
func shortestPathLength(graph [][]int) int {
	sead := make(map[int]*VNode)
	for i := range graph {
		appendNode(i, &graph, &sead)
	}
	maxLength := 0
	for _, n := range sead {
		if n.MaxLength > maxLength {
			maxLength = n.MaxLength
		}
	}
	return maxLength - 1 + 2*(len(graph)-maxLength)
}

func appendNode(i int, graph *[][]int, sead *map[int]*VNode) {
	n := VNode{
		Val:       i,
		MaxLength: 1,
		Link:      []int{},
	}
	for _, v := range (*graph)[i] {
		if _, ok := (*sead)[v]; !ok {
			continue
		}
		t := (*sead)[v]
		n.MaxLength = t.MaxLength + 1
		n.Link = append(n.Link, v)
		t.Link = append(t.Link, i)
	}
}

// @lc code=end
