package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//97.59%/30.25%

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1 := headA
	cur2 := headB
	for {
		if cur1 == nil && cur2 == nil {
			return nil
		} else if cur1 == nil {
			cur2 = cur2.Next
			cur1 = headB
		} else if cur2 == nil {
			cur1 = cur1.Next
			cur2 = headA
		} else if cur1 == cur2 {
			return cur1
		} else {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
	}
}
