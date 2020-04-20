package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	para string
}

type answer struct {
	ans int
}

type question struct {
	p param
	a answer
}

func TestLengthOfLongestSubstring(t *testing.T) {
	// 断言
	ast := assert.New(t)

	// 样例
	example := []question{
		{
			p: param{
				"abcabcbb",
			},
			a: answer{
				3,
			},
		},
		{
			p: param{
				"bbbbb",
			},
			a: answer{
				1,
			},
		},
		{
			p: param{
				"pwwkew",
			},
			a: answer{
				3,
			},
		},
	}

	for _, qu := range example {
		ans, para := qu.a, qu.p
		ast.Equal(ans.ans, lengthOfLongestSubstring(para.para), "样例输入：%v", para)
	}
}
