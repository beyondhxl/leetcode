/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0003

func lengthOfLongestSubstring(s string) int {
	// 思路一
	// 遍历字符串，将不重复的子串都保存下来，然后比较长度
	// 两个游标要同时移动，比对的是相邻两个字符，而不是一直和第一个字符做比较
	/*
		var str, substr []byte
		str = []byte(s)
		count := 0
		m := make(map[int][]byte) // 长度与子字符串的映射
		for i := 0; i < len(str); {
			for j := i + 1; j < len(str); {
				if str[i] != str[j] {
					count++
					if len(substr) == 0 {
						substr = append(substr, str[i])
					}
					substr = append(substr, str[j])
					i++
					j++
				} else {
					m[count] = substr
					count = 0
					substr = nil
					i++
					j++
					break
				}
			}
		}
		maxlen := 0
		for k, v := range m {
			if k > maxlen {
				maxlen = k
			}
			fmt.Print(v)
		}
		return maxlen
	*/
	var (
		start        = 0
		maxlength    = 0
		charlocation = make(map[rune]int) // 缓存每次出现的字符及位置索引
	)

	for i, ch := range []rune(s) {
		if lastidx, ok := charlocation[ch]; ok && lastidx >= start {
			start = lastidx + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		charlocation[ch] = i
	}
	return maxlength
}
