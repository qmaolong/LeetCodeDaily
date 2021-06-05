package main

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	cursor := head
	for cursor != nil {
		if head.Val == val {
			head = head.Next
			cursor = head
			continue
		}
		next := cursor.Next
		for next != nil && next.Val == val {
			cursor.Next = next.Next
			next = cursor.Next
		}
		cursor = cursor.Next
	}
	return head
}

// @lc code=end
