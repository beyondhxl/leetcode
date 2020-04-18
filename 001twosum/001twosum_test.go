package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Error("Failed, Not Find")
	}
}
