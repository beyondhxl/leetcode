/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0006

import "bytes"

/*
	以numRows=3为例，行优先打印出字符

 		行数	0|		1   5   9     13
 			1|		2 4 6 8 10 12 14 16
 			2|		3   7   11    15

	每一行右边的字符的索引值都是其左边的字符的索引值加上它下面剩余行数的两倍或上面行数的两倍(具体看两个字符所在的行)
	因为两个字符目前在同一行，所以左边的字符是经过从上到下的，然后斜着从左到右到右边的字符，或者相反。例如：
			4 == 2(左边的索引) + 2(两倍) * 1(4的下面有一行)
			7 == 3(左边的索引) + 2(两倍) * 2(3的上面有两行)
*/

func convert(s string, numRows int) string {
	return convert1(s, numRows)
	return convert2(s, numRows)
}

func convert1(s string, numRows int) string {
	if numRows <= 1 || len(s) <= numRows {
		return s
	}

	var res []byte
	for i := 0; i < numRows; i++ {
		count := 0                  // 计数
		upline := i                 // 上方的行数
		downline := numRows - 1 - i // 下方的行数
		for j := i; j < len(s); {
			if count%2 == 0 && downline != 0 {
				res = append(res, s[j])
				j += 2 * downline
			} else if count%2 != 0 && upline != 0 {
				res = append(res, s[j])
				j += 2 * upline
			}
			count++
		}
	}
	return string(res)
}

func convert2(s string, numRows int) string {
	// 从上到下、然后从左到右遍历字符串
	// 需要算出有多少列？
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// 步长
	// 因为计算两个位置索引差时，没有包含当前行，所以是2*(numRows-1)==numRows*2-2
	step := numRows*2 - 2

	// 首行
	for i := 0; i < len(s); i += step {
		res.WriteByte(s[i])
	}

	// 中间行
	for j := 1; j <= numRows-2; j++ {
		// 每行的第一个元素
		res.WriteByte(s[j])

		for k := step; k-j < len(s); k += step {
			res.WriteByte(s[k-j])
			if k+j < len(s) {
				res.WriteByte(s[k+j])
			}
		}
	}

	// 末行
	for l := numRows - 1; l < len(s); l += step {
		res.WriteByte(s[l])
	}

	return res.String()
}
