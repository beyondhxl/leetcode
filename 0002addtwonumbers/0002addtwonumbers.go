/*

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Val值逆序存储
	// 要考虑进位
	/*
		res := &ListNode{}
		cur := res
		for l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
			if l1.Val+l2.Val <= 0 {
				continue
			} else if l1.Val+l2.Val < 10 {
				cur.Val = l1.Val + l2.Val
			} else {
				cur.Val = 0
				cur.Next = &ListNode{}
				cur.Next.Val += 1
			}
			l1 = l1.Next
			l2 = l2.Next
		}
		return res.Next
	*/
	res := &ListNode{}
	cur := res
	carry := 0 // 进位
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10
		lnode := &ListNode{Val: val}
		cur.Next = lnode
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	// 最后无进位
	if carry == 0 {
		// l1比l2长
		if l1 != nil {
			cur.Next = l1
		} else {
			cur.Next = l2
		}
	} else {
		for l1 != nil {
			sum := l1.Val + carry
			val := sum % 10
			carry = sum / 10
			node := &ListNode{Val: val}
			cur.Next = node
			cur = cur.Next
			l1 = l1.Next
		}
		for l2 != nil {
			sum := l2.Val + carry
			val := sum % 10
			carry = sum / 10
			node := &ListNode{Val: val}
			cur.Next = node
			cur = cur.Next
			l2 = l2.Next
		}
		for carry != 0 {
			val := carry % 10
			carry = carry / 10
			node := &ListNode{Val: val}
			cur.Next = node
			cur = cur.Next
		}
	}
	return res.Next
}
