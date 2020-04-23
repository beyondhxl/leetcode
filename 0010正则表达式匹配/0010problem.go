/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0010

/*
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
*/
func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)

	dp := make([][]bool, slen+1)
	for i := range dp {
		dp[i] = make([]bool, plen+1)
	}
	// dp[i][j]代表s[:i]能否与p[:j]匹配
	dp[0][0] = true

	for j := 1; j < plen && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < slen; i++ {
		for j := 0; j < plen; j++ {
			/*
				p[j]可以匹配任意单个字符或者与当前等待匹配的字符相等
			*/
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/*
					p[j]的匹配情况与它的前一个字符p[j-1]的内容相关
				*/
				if p[j-1] != s[i] && p[j-1] != '.' {
					/*
						p[j]与s[i]匹配不上，p[j-1:j+1]只能被当做''空字符串
					*/
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/*
						p[j]与s[i]匹配上，p[j-1:j+1]作为 'x*' 有如下三种解释
					*/
					dp[i+1][j+1] = dp[i+1][j-1] || // 'x*' 解释为 ''
						dp[i+1][j] || // 'x*' 解释为 'x'
						dp[i][j+1] // 'x*' 解释为 'xx...'
				}
			}
		}
	}

	return dp[slen][plen]
}
