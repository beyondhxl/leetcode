/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0022

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(l, r, d int, bytes []byte, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, string(bytes))
		return
	}

	if l > 0 {
		bytes[d] = '('
		dfs(l-1, r, d+1, bytes, res)
	}

	if r > 0 && l < r {
		bytes[d] = ')'
		dfs(l, r-1, d+1, bytes, res)
	}
}
