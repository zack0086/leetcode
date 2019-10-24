/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var pre *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	pre = head
	for {
		if l1 == nil {
			pre.Next = l2
			break
		}
		if l2 == nil {
			pre.Next = l1
			break
		}

		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	return head
}

// @lc code=end

