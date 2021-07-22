package main

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//  19/19 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 7.95 % of golang submissions (3.7 MB)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	return copyNode(head, m)
}

func copyNode(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	} else if v, ok := m[node]; ok {
		return v
	}
	n := &Node{
		Val: node.Val,
	}
	m[node] = n
	n.Next = copyNode(node.Next, m)
	n.Random = copyNode(node.Random, m)
	return n
}

// @lc code=end

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
