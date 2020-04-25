/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/*
		if head == nil || n == 0 {
			return nil
		}

		count := 0
		for head != nil {
			count++
			head = head.Next
		}
		if n > count {
			return nil
		}

		front := &ListNode{}
		for head != nil {
			n--
			if n == 0 {
				front = head
			}
			head = head.Next
		}

		for front != nil {
			front = front.Next
			head = head.Next
			return head
		}
		return nil
	*/
	dummyhead := &ListNode{}
	dummyhead.Next = head

	p := dummyhead
	q := dummyhead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return dummyhead.Next
}
