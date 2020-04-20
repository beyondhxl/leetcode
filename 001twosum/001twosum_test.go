package main

import (
	"fmt"
	"reflect"
	"testing"
)

var nums []int = []int{2, 7, 11, 15}

func TestTwoSum(t *testing.T) {
	res := twoSum(nums, 9)
	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Error("Failed, Not Find")
	}

}

func TestTwoSum1(t *testing.T) {
	res1 := twoSum1(nums, 9)
	fmt.Print(res1)
	if !reflect.DeepEqual(res1, []int{0, 1}) {
		t.Error("Failed, Not Find")
	}
}
