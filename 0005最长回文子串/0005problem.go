/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0005

// 正反读都一样的字符串
// 上海自来水来自海上
// aba
func longestPalindrome(s string) string {
	// 只有一个字符
	if len(s) < 2 {
		return s
	}

	// 最长回文起始位置、最长回文的长度
	start, maxlen := 0, 0

	for i := 0; i < len(s); {
		if len(s)-i <= maxlen/2 {
			/*
				i是回文字符串正中间部分首字符的索引
				假设maxlen=1，则 1 <= (len(s)-i)*2 - 1，因为i的坐标范围是0~(len(s)/2 -1)
				若 len(s)-i <= maxlen/2，则 1 <= maxlen-1，进一步 1 < maxlen，与假设矛盾
			*/
			break
		}

		// beg：回文首字符索引
		// end：回文尾字符索引
		beg, end := i, i
		// 如果有与中间字符相同的字符，则继续向右边扩展
		for end < len(s)-1 && s[end+1] == s[end] {
			end++
			// 循环结束后，s[beg:end+1]是一串相同的字符
		}

		// 下一次循环，下一个回文的正中间字符串的首字符是s[end+1]
		// 开始向左边扩展
		i = end + 1
		for end < len(s)-1 && beg > 0 && s[end+1] == s[beg-1] {
			end++
			beg--
			// 循环结束后，s[beg:end+1]就是找到的最大的回文字符串
		}

		newlen := end + 1 - beg
		if newlen > maxlen {
			start = beg
			maxlen = newlen
		}
	}
	return s[start : start+maxlen]
}
