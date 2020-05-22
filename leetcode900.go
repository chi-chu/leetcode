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
	lens := len(nums)
	if lens < 1 {
		return 0
	}
	QuickSort(nums, 0, lens-1)
	fmt.Println(nums)
	val := nums[0]
	index := 0
	for k,v := range nums {
		//fmt.Println((val - v) - (k - index))
		if val + (k-index) - v >= 0 {
			ret += val + (k-index) - v
		} else{
			val = v
			index = k
		}
		//fmt.Println(k, ret)
	}
	return ret
}

func leetcode994(grid [][]int) int {
	var time, good int
	var location [][2]int
	for j,v := range grid {
		for i,vv := range v {
			if vv == 1 {
				good++
			}else if vv == 2 {
				location = append(location, [2]int{i,j})
			}
		}
	}
	//fmt.Println(location)
	for {
		if good == 0 {
			return time
		}
		if len(location) == 0 {
			break
		}
		time++
		tmp := make([][2]int,0)
		for _,v := range location {
			if v[0]+1 < len(grid[0]) && grid[v[1]][v[0]+1] == 1 {
				grid[v[1]][v[0]+1] = 2
				good--
				tmp = append(tmp, [2]int{v[0]+1, v[1]})
			}
			if v[0]-1 >= 0 && grid[v[1]][v[0]-1] == 1 {
				grid[v[1]][v[0]-1] = 2
				good--
				tmp = append(tmp, [2]int{v[0]-1, v[1]})
			}
			if v[1]+1 < len(grid) && grid[v[1]+1][v[0]] == 1 {
				grid[v[1]+1][v[0]] = 2
				good--
				tmp = append(tmp, [2]int{v[0], v[1]+1})
			}
			if v[1]-1 >= 0 && grid[v[1]-1][v[0]] == 1 {
				grid[v[1]-1][v[0]] = 2
				good--
				tmp = append(tmp, [2]int{v[0], v[1]-1})
			}
		}
		location = tmp
	}
	if good > 0 {
		return -1
	}
	return time-1
}