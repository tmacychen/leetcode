package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	l := len(nums)
	tmpMap := make(map[int]int, len(nums))
	for i := 0; i < l; i++ {
		tmpMap[nums[i]] = i
	}

	for i := 0; i < l; i++ {
		index, ok := tmpMap[target-nums[i]]
		if ok && i != index {
			return []int{i, index}
		}
	}
	return nil
}

func main() {
	//	nums := []int{1, 2, 3, 4, 5}
	//target := 8
	nums := []int{11, 15, 1, 8, 0, 9}
	target := 9

	ret := twoSum(nums, target)
	fmt.Printf("ret = %v\n", ret)
}
