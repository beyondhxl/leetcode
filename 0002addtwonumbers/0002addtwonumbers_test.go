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

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	res := &ListNode{
		Val: vals[0],
	}
	cur := res
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{
			Val: vals[i],
		}
		cur = cur.Next
	}
	return res
}

func printList(l *ListNode) {
	for l != nil {
		val := l.Val
		l = l.Next
		if l == nil {
			fmt.Printf("%d", val)
		} else {
			fmt.Printf("%d --> ", val)
		}
	}
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	ans := createList([]int{7, 0, 8})
	res := addTwoNumbers(l1, l2)
	ast.Equal(ans, res)
	fmt.Printf("输入：(")
	printList(l1)
	fmt.Printf(") + (")
	printList(l2)
	fmt.Println(")")
	fmt.Printf("输出：")
	printList(res)
	fmt.Println()
}
