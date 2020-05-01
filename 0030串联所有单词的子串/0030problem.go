/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。



示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0030

// 解法一 待研究
func findSubstring(s string, words []string) []int {
	lens := len(s)
	res := make([]int, 0, lens)

	lenws := len(words)
	if lens == 0 || lenws == 0 || lens < lenws*len(words[0]) {
		return res
	}

	lenw := len(words[0])

	// record记录words中每个单词出现的次数
	record := make(map[string]int, lenws)
	for _, w := range words {
		if len(w) != lenw {
			return res
		}
		record[w]++
	}

	// remain记录words中每个单词还能出现的次数
	remain := make(map[string]int, lenws)
	// count记录符合要求的单词的连续出现次数
	count := 1
	left, right := 0, 0

	// reset重置remain和count
	reset := func() {
		if count == 0 {
			return
		}
		for k, v := range record {
			remain[k] = v
		}
		count = 0
	}

	// moveLeft让left指向下一个单词
	moveLeft := func() {
		// 增加left指向word可出现次数一次
		remain[s[left:left+lenw]]++
		// 连续符合条件的单词数减少一次
		count--
		// left后移一个word的长度
		left += lenw
	}

	// left需要分别从{0，1，... ，lenw-1}这些位置开始检查
	for i := 0; i < lenw; i++ {
		left, right = i, i
		reset()

		// s[left:]的长度大于words中所有word组成的字符串的长度时，
		// s[left:]中才有可能存在要找的字符串
		for lens-left >= lenws*lenw {
			word := s[right : right+lenw]
			remainTimes, ok := remain[word]

			switch {
			case !ok:
				// word不在words中
				// 从right+lenw处，作为新窗口，重新开始统计
				left, right = right+lenw, right+lenw
				reset()
			case remainTimes == 0:
				// word的出现次数上次就用完了
				// 说明word在s[left:right]中出现次数过多
				// 这个case会连续出现
				// 直到s[left:right]中的统计结果是remain[word]==1
				// 这个过程中，right一直不变
				moveLeft()
			default:
				// ok && remainTimes > 0，word符合出现的条件
				// moveRight
				remain[word]--
				count++
				right += lenw
				// 检查words能否排列组合成s[left:right]
				if count == lenws {
					res = append(res, left)
					// moveLeft可以避免重复统计s[left+lenw:right]中的信息
					moveLeft()
				}
			}
		}
	}
	return res
}

// 解法二 一种朴素的匹配方法，每次一个字母一个字母地移动
func findSubstring1(s string, words []string) []int {
	var res []int
	wordnum := len(words)
	if wordnum == 0 {
		return res
	}
	wordlen := len(words[0])
	allwordmap := make(map[string]int)
	for _, w := range words {
		if len(w) != wordlen {
			return res
		}
		allwordmap[w]++
	}
	lens := len(s)
	for i := 0; i < lens-wordnum*wordlen+1; i++ {
		haswordmap := make(map[string]int)
		// 判断子串是否符合
		num := 0
		for num < wordnum {
			word := s[i+num*wordlen : i+(num+1)*wordlen]
			if _, ok := allwordmap[word]; ok {
				haswordmap[word]++
				if haswordmap[word] > allwordmap[word] {
					break
				}
			} else {
				break
			}
			num++
		}
		if num == wordnum {
			res = append(res, i)
		}
	}
	return res
}
