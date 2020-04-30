/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0023

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	llen := len(lists)
	mid := llen / 2

	if llen == 1 {
		return lists[0]
	}

	if llen == 2 {
		l1, l2 := lists[0], lists[1]
		res, cur := &ListNode{}, &ListNode{}

		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}

		if l1.Val < l2.Val {
			res = l1
			cur.Next = l1
			l1 = l1.Next
		} else {
			res = l2
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		}

		if l1 != nil {
			cur.Next = l1
		}

		if l2 != nil {
			cur.Next = l2
		}

		return res
	}

	return mergeKLists([]*ListNode{mergeKLists(lists[mid:]), mergeKLists(lists[:mid])})
}

// 解法二
// 优先队列
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	// golang容器中的heap
	var h MinHeap
	heap.Init(&h)

	for _, l := range lists {
		if l != nil {
			heap.Push(&h, l)
		}
	}

	dummy := &ListNode{}
	p := dummy

	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	num := len(old)
	x := old[num-1]
	*h = old[0 : num-1]
	return x
}
