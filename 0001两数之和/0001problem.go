package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。


示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/

// 两个for循环遍历，时间复杂度较高O(n^2)
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0)
	nlen := len(nums)
	for i := 0; i < nlen; i++ {
		for j := i + 1; j < nlen; j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
			}
		}
	}
	return ret
}

// 用一个map存储遍历过程中出现的数字，key就是数字
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i} // 第一个数已经在之前出现了，此时第二个数的下标肯定是在后面的
		}
		m[v] = i // 如果不在map中，则将自己与对应自己的下标存入
	}
	return nil
}
