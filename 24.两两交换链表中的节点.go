/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	falseHead := &ListNode{Val:0, Next:head}
	pre := falseHead
	node := head
	for node != nil && node.Next != nil{
		pre.Next = node.Next
		node.Next = node.Next.Next
		pre.Next.Next = node
		pre = node
		node = node.Next
	}
	return falseHead.Next
}
// @lc code=end

