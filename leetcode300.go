package main

import (
	"math"
	"strings"
)

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

/*
预备知识：贝祖定理
我们认为，每次操作只会让桶里的水总量增加 x，增加 y，减少 x，或者减少 y。
你可能认为这有问题：如果往一个不满的桶里放水，或者把它排空呢？那变化量不就不是 x 或者 y 了吗？接下来我们来解释这一点：
首先要清楚，在题目所给的操作下，两个桶不可能同时有水且不满。因为观察所有题目中的操作，操作的结果都至少有一个桶是空的或者满的；
其次，对一个不满的桶加水是没有意义的。因为如果另一个桶是空的，那么这个操作的结果等价于直接从初始状态给这个桶加满水；而如果另一个桶是满的，
那么这个操作的结果等价于从初始状态分别给两个桶加满；再次，把一个不满的桶里面的水倒掉是没有意义的。因为如果另一个桶是空的，
那么这个操作的结果等价于回到初始状态；而如果另一个桶是满的，那么这个操作的结果等价于从初始状态直接给另一个桶倒满。
因此，我们可以认为每次操作只会给水的总量带来 x 或者 y 的变化量。因此我们的目标可以改写成：找到一对整数 a, ba,b，使得
ax+by=z
而只要满足 z≤x+y，且这样的 a, b存在，那么我们的目标就是可以达成的。这是因为：
若 a≥0,b≥0，那么显然可以达成目标。
若 a<0，那么可以进行以下操作：
往 y 壶倒水；
把 y 壶的水倒入 x 壶；
如果 y 壶不为空，那么 x 壶肯定是满的，把 x 壶倒空，然后再把 y 壶的水倒入 x 壶。
重复以上操作直至某一步时 x 壶进行了 aa 次倒空操作，y 壶进行了 bb 次倒水操作。
若 b<0，方法同上，x 与 y 互换。
而贝祖定理告诉我们，ax+by=z 有解当且仅当 zz 是 x, yx,y 的最大公约数的倍数。因此我们只需要找到 x, yx,y 的最大公约数并判断 zz 是否是它的倍数即可。
 */
func leetcode365(x, y, z int) bool {
	if z == 0 {
		return true
	}
	if x == 0 || y == 0 {
		return z == x || z == y
	}
	g := GCD(x, y)
	return z%g == 0 && z >= 0 && z <= (x+y)
}

func GCD(x, y int) int {
	p1, p2 := x, y
	if p2 > p1 {
		p1, p2 = y, x
	}
	if p1%p2 == 0 {
		return p2
	}
	return GCD(p1%p2, p2)
}

func leetcode394(s string) string {
	var ret string
	var copyItem string
	var repeatStart int
	var repeatEnd int
	var nums int
	for _, v := range s {
		if v > '0' && v < '9' {
			nums = nums*10 + int(v - '0')
		} else if v == '[' {
			repeatStart++
		} else if v == ']' {
			repeatEnd++
			if repeatStart == repeatEnd {
				ret += strings.Repeat(copyItem, nums)
				nums = 0
				copyItem = ""
			}else{
				copyItem = strings.Repeat(copyItem, nums)
			}
		} else {
			if repeatStart > repeatEnd {
				copyItem += string(v)
			}else{
				ret += string(v)
			}
		}
	}
	return ret
}