package main

import (
	"fmt"
)

func leetcode912(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func leetcode914(deck []int) bool {
	gcd := -1
	var first int
	var flag bool
	scount := make(map[int]int)
	for _,v := range deck {
		_, ok := scount[v]
		if ok {
			scount[v]++
		}else{
			scount[v] = 1
		}
	}
	for _,v := range scount {
		if !flag {
			first = v
			flag = true
		}else{
			if gcd == -1 {
				gcd = GCD(first, v)
				if gcd < 2 {
					return false
				}
			}else{
				if gcd != GCD(gcd, v) {
					return false
				}
			}
		}
	}
	if len(scount) == 1 && first < 2 {
		return false
	}
	return true
}
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
