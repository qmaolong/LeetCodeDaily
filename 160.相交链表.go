package main

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  时间复杂度2（max(m,n)）,空间复杂度4
//	可优化至时间复杂度m+n，空间复杂度2
//	思路：无需计算链表长度，利用m+n=n+m达到步伐一致
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	var aLen, bLen int
	aCursor := headA
	bCursor := headB
	for {
		needBreak := true
		if aCursor != nil {
			aLen++
			aCursor = aCursor.Next
			needBreak = false
		}
		if bCursor != nil {
			bLen++
			bCursor = bCursor.Next
			needBreak = false
		}
		if needBreak {
			break
		}
	}
	aCursor = headA
	bCursor = headB
	for {
		if aCursor == bCursor {
			return aCursor
		}
		if aLen > bLen {
			aCursor = aCursor.Next
			aLen--
		} else if aLen < bLen {
			bCursor = bCursor.Next
			bLen--
		} else {
			aCursor = aCursor.Next
			aLen--
			bCursor = bCursor.Next
			bLen--
		}
		if aLen == 0 || bLen == 0 {
			break
		}
	}
	return nil
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
