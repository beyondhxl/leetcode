/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。


示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0010

func maxArea(height []int) int {
	// 可以转换为二维数组
	i, j := 0, len(height)-1
	res := 0

	for i < j {
		h1, h2 := height[i], height[j]
		h := minnum(h1, h2)

		area := h * (j - i)
		if res < area {
			res = area
		}

		if h1 < h2 {
			i++
		} else {
			j--
		}
	}

	return res
}

func minnum(l, r int) int {
	if l <= r {
		return l
	}
	return r
}
