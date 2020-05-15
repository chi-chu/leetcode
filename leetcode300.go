package main

import "math"

func leetcode300(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ret int
	var tmp []int
	tmp = append(tmp, nums[0])
	for _,v := range nums {
		//fmt.Println(tmp,"---", v)
		if v > tmp[ret] {
			tmp = append(tmp, v)
			ret++
		}else{
			var left, right int
			right = ret
			for left < right {
				mid := (right - left) >>1 + left
				if tmp[mid] < v {
					//只有小于 才 +1
					left = mid + 1
				} else {
					right = mid
				}
			}
			tmp[left] = v
		}
	}
	return ret + 1
}

//背包问题
func leetcode322(coins []int, amount int) int {
	dp := make([]int, amount + 1, amount+1)
	for k,_ := range dp {
		//max +1  = -max
		dp[k] = math.MaxInt64-1
	}
	dp[0] = 0
	for _,coin := range coins {
		for i:=coin; i<amount+1; i++ {
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin]+1
			}
		}
	}
	return dp[amount]
}

func leetcode344(s []byte) {
	var index int
	var tmp byte
	for len(s) - 1 - index > index {
		tmp = s[index]
		s[index] = s[len(s)-1-index]
		s[len(s)-1-index] = tmp
		index++
	}
}