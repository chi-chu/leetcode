package main

import "fmt"

func leetcode945(nums []int) int {
	var ret int
	if nums ==  nil {
		return 0
	}
	lens := len(nums)
	QuickSort(nums, 0, lens-1)
	fmt.Println(nums)
	for k,v := range nums {
		if nums[0] + k - v >= 0 {
			ret += nums[0] + k - v
		} else{
			break
		}
	}
	return ret
}
