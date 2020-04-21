/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// O(log(m+n))的时间复杂度
	// 先合并两个数组
	nums := mergeArray(nums1, nums2)

	// 返回合并后的中位数
	return median(nums)
}

func mergeArray(nums1 []int, nums2 []int) []int {
	var i, j int
	len1 := len(nums1)
	len2 := len(nums2)
	res := make([]int, len1+len2)
	lenr := len(res)
	// 合并都是先将两个数组中数值小的加入到结果数组中
	for k := 0; k < lenr; k++ {
		// 数组1走到末尾
		if i == len1 || (i < len1 && j < len2 && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j++
			continue
		}

		// 数组2走到末尾
		if j == len2 || (i < len1 && j < len2 && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i++
		}
	}
	return res
}

func median(nums []int) float64 {
	l := len(nums)
	if l == 0 {
		panic("切片长度异常")
	}
	// 根据切片的长度的数值的奇偶区分
	if l%2 == 0 {
		// 最中间的两个数取平均
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	} else {
		// 最中间的数/2
		return float64(nums[l/2])
	}
}
