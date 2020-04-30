/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0024

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一
// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// 解法二
// 非递归
func swapPairs1(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	cur := pre

	for cur.Next != nil && cur.Next.Next != nil {
		start := cur.Next
		end := cur.Next.Next
		cur.Next = end
		start.Next = end.Next
		end.Next = start
		cur = start
	}
	return pre.Next
}
