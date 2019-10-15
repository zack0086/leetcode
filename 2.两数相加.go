/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := l1
	up := 0
	for {
		temp := l1.Val + l2.Val + up
		if temp > 9 {
			temp -= 10
			up = 1
		} else {
			up = 0
		}
		l1.Val = temp
		if l1.Next == nil || l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1.Next == nil && l2.Next == nil{
		if up == 1{
			node := &ListNode{Val: 1, Next: nil}
			l1.Next = node
		}
		return ans
	} else if l1.Next == nil {
		l1.Next = l2.Next
	}
	l1 = l1.Next
	for up != 0 {
		temp := l1.Val + up
		if temp > 9 {
			temp -= 10
			up = 1
		} else {
			up = 0
		}
		l1.Val = temp
		if l1.Next == nil {
			if up == 1{
				node := &ListNode{Val: 1, Next: nil}
				l1.Next = node
			}
			break
		}
		l1 = l1.Next
	}
	return ans
}
// @lc code=end

