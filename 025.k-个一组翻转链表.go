/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
	falseHead := &ListNode{Val:0, Next:head}
	lastTail := falseHead
	cur := head
	for {
		if ifRev(cur, k) {
			lastTail = rev(lastTail, k)
			cur = lastTail.Next
		} else {
			break
		}
	}
	return falseHead.Next
}

func rev(lastTail *ListNode, k int) *ListNode {
	tmp := lastTail.Next
	node := lastTail.Next
	pre := lastTail
	for k > 0 {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
		k--
	}
	lastTail.Next.Next = node
	lastTail.Next = pre
	return tmp
}

func ifRev(head *ListNode, k int) bool {
	for k > 0 {
		if head == nil {
			return false
		}
		head = head.Next
		k--
	}
	return true
}
// @lc code=end

