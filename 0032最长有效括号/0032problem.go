/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0032

// 解法一 DP
// 分析 https://leetcode-cn.com/problems/longest-valid-parentheses/solution/dong-tai-gui-hua-si-lu-xiang-jie-c-by-zhanganan042/
func longestValidParentheses(s string) int {
	longestlen := 0
	if len(s) < 2 {
		return longestlen
	}

	dp := make([]int, len(s))
	dp[0] = 0
	if s[1] == '(' {
		dp[1] = 0
	} else if s[0] == '(' {
		dp[1] = 2
		longestlen = 2
	}

	for i := 2; i < len(s); i++ {
		// 最长有效括号必定是以')'结尾
		if s[i] == '(' {
			dp[i] = 0
			continue
		}
		if s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else {
			// 左边想要配对的'('的索引不合法，都越界了
			if i-1-dp[i-1] < 0 {
				dp[i] = 0
			} else if s[i-1-dp[i-1]] == '(' {
				if i-1-dp[i-1]-1 < 0 {
					dp[i] = dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]
				}
			}
		}
		if dp[i] > longestlen {
			longestlen = dp[i]
		}
	}
	return longestlen
}

// TODO 解法二 不需要额外的空间
