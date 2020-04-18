package main

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
