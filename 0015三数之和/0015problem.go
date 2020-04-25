/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 左右指针指向nums[i]后面的数组两端
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 去重
				for l < r && nums[l] == nums[r] {
					if nums[l] == nums[l+1] {
						l++
					}
				}
				for l < r && nums[r] == nums[r-1] {
					if nums[r] == nums[r-1] {
						r--
					}
				}
				l++
				r--
			}
		}
	}
	return res
}
