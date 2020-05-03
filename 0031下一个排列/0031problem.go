/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0031

import "sort"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 找到从右向左的第一个nums[i] < nums[j]
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 从nums[j:]中找到第一个大于nums[i]的
	if i >= 0 {
		k := len(nums) - 1
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 将nums[j]升序排序
	sort.Ints(nums[j:])
}
