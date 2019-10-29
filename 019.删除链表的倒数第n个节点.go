/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ans := head
	if head == nil {
		return head
	}
	node := head
	var preNode *ListNode
	for node != nil {
		next := node.Next
		node.Next = preNode
		preNode = node
		node = next
	}
	flag := 1
	node = preNode
	preNode = nil
	for node != nil {
		if flag == n {
			node = node.Next
			if node == nil {
				ans = preNode
				break
			}
		}
		next := node.Next
		node.Next = preNode
		preNode = node
		node = next
		flag++
	}
	return ans
}

// @lc code=end

