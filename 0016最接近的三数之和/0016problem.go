/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// 排序，排序后数组元素是递增的，方便后面指针的移动
	if len(nums) < 3 {
		panic("数组个数不对")
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[l] + nums[r] + nums[i]
			if s > target {
				r--
			} else if s < target {
				l++
			} else {
				return s
			}
			if math.Abs(float64(target-s)) < math.Abs(float64(target-res)) {
				res = s
			}
		}
	}
	return res
}
