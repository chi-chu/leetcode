package main

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