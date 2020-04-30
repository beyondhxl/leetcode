/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	end := dummy

	for {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		start := pre.Next
		next := end.Next
		pre.Next = reverse(start, end.Next)
		start.Next = next
		pre = start
		end = start
	}
	return dummy.Next
}

func reverse(head, end *ListNode) *ListNode {
	pre := &ListNode{}
	for head != end {
		/*
			head.Next = pre
			pre = head
			head = head.Next
		*/
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	left := dummy // 第二个指针
	right := head // 第一个移动的指针
	move := false // left指针移动标志
	count := 0    // 移动计数
	for right != nil {
		if count == k {
			count = 0
			move = true
			left.Next = reverse(left.Next, k)
		}
		right = right.Next
		count++
		if move {
			left = left.Next
		}
	}
	if count == k {
		left.Next = reverse(left.Next, k)
	}
	return dummy.Next
}

func reverse1(head *ListNode, k int) *ListNode {
	count := 1
	p := head.Next
	first := head
	var two *ListNode
	for count < k {
		two = p
		if p != nil {
			p = p.Next
		} else {
			break
		}
		count++
		two.Next = first
		first = two
	}
	head.Next = p
	return two
}
