/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem009

import "strconv"

// 解法一
// 反转一半数字
func isPalindrome1(x int) bool {
	// 负数不是回文数
	// 如果数字最后一位是0，只有数字0满足回文特征，因为多位整数不可能是0开头的
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertnum := 0
	for x > revertnum {
		revertnum = revertnum*10 + x%10
		x = x / 10
	}

	// 数字长度是奇数时，中间的数字不影响回文的，如12321
	return x == revertnum || x == revertnum/10
}

//	解法二
//	转换为字符串
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
