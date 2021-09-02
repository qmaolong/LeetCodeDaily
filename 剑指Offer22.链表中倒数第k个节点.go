package main

/*
 *
 * 剑指 Offer 22. 链表中倒数第k个节点
 */

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00%的用户
// 通过测试用例：208 / 208

//  * Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func getKthFromEnd(head *ListNode, k int) *ListNode {
	c1 := head
	c2 := head
	for c2 != nil {
		if k > 0 {
			c2 = c2.Next
			k--
		} else {
			c1 = c1.Next
			c2 = c2.Next
		}
	}
	return c1
}
