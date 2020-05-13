package main

import (
	"fmt"
	"math"
)

func leetcode102(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var view [][]int
	echoTree(&view, root, 0)
	return view
}

func echoTree(view *[][]int, root *TreeNode, level int) {
	if len(*view) - 1 < level {
		*view = append(*view, []int{})
	}
	(*view)[level] = append((*view)[level], root.Val)
	if root.Left != nil {
		echoTree(view, root.Left, level+1)
	}
	if root.Left == nil && root.Right != nil {
		if len(*view) - 2 < level {
			*view = append(*view, []int{})
		}
	}
	if root.Right != nil {
		echoTree(view, root.Right, level+1)
	}
}

func leetcode121(prices []int) int {
	var ret, preMin int
	preMin = math.MaxInt64
	for _,v := range prices {
		if v < preMin {
			preMin = v
		}else{
			if v - preMin > ret {
				ret = v - preMin
			}
		}
	}
	return ret
}

func leetcode136(nums []int) int {
	//交换律：a ^ b ^ c <=> a ^ c ^ b
	//任何数于0异或为任何数 0 ^ n => n
	//相同的数异或为0: n ^ n => 0
	var tmp int
	for _,v := range nums {
		tmp ^= v
	}
	return tmp
}


type LRUCache struct {
	cap		int
	now		int
	keys	*Node
	values	map[int]int
}

//下述结构效率高很多
/*type LinkNode struct {
	key,value int
	next,prev *LinkNode
}

type LRUCache struct {
	m map[int] *LinkNode
	cap int
	head,tail *LinkNode
}*/

//执行结果：通过 显示详情
//执行用时 :304 ms, 在所有 Go 提交中击败了5.13%的用户
//内存消耗 :12 MB, 在所有 Go 提交中击败了100.00%的用户  想哭。。。
func leetcode146() {
	l := Constructor(2)
	l.Put(2, 1)
	fmt.Println(l.Get(2))
	l.Put(3,2)
	l.keys.Print()
	fmt.Println(l.Get(2))
	l.keys.Print()
	fmt.Println(l.Get(3))
}

func Constructor(capacity int) LRUCache {
	o := LRUCache{}
	o.cap = capacity
	o.values = make(map[int]int)
	return o
}

func (l *LRUCache) Get(key int) int {
	o,ok := l.values[key]
	if !ok {
		return -1
	}
	tmp := l.keys
	if tmp.Value == key {
		return o
	}
	var before, next *Node
	for tmp.Next != nil {
		if tmp.Value == key {
			next = tmp.Next
			break
		}
		before = tmp
		tmp = tmp.Next
	}
	if tmp != nil {
		tmp.Next = l.keys
		if before != nil {
			before.Next = next
		}
		l.keys = tmp
	}
	return o
}

func (l *LRUCache) Put(key int, value int) {
	_, ok := l.values[key]
	l.values[key] = value
	if ok {
		tmp := l.keys
		var before, next *Node
		if tmp.Value == key {
			return
		}
		for tmp != nil {
			if tmp.Value == key {
				next = tmp.Next
				break
			}
			before = tmp
			tmp = tmp.Next
		}
		if tmp != nil {
			tmp.Next = l.keys
			if before != nil {
				before.Next = next
			}
			l.keys = tmp
		}
	}else{
		l.now++
		l.keys = &Node{key, l.keys}
		if l.now > l.cap {
			l.now = l.cap
			tmp := l.keys
			for i:=0; i<l.cap-1; i++{
				tmp = tmp.Next
			}
			c := tmp.Next
			for c != nil {
				delete(l.values, c.Value)
				c = c.Next
			}
			tmp.Next = nil
		}
	}
}

//动态规划 典型 dp[i] = max(dp[i-1], num[i]+ dp[i-2])
func leetcode198(nums []int) int {
	var ret int
	lens := len(nums)
	if lens == 0{
		return 0
	}else if lens == 1 {
		return nums[0]
	}else if lens == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}else{
			return nums[1]
		}
	}else{
		dp := make([]int, lens, lens)
		dp[0] = nums[0]
		if nums[0] > nums[1] {
			dp[1] = nums[0]
			ret = nums[0]
		}else{
			dp[1] = nums[1]
			ret = nums[1]
		}
		for i:=2; i<lens; i++ {
			if dp[i-2] + nums[i] > dp[i-1] {
				dp[i] = nums[i] + dp[i-2]
				ret = dp[i]
			}else{
				dp[i] = dp[i-1]
			}
		}
	}
	return ret
}

func leetcode199(root *TreeNode) []int {
	var ret, indexCompare []int
	if root == nil {
		return nil
	}
	rightSideView(root, 1, 1, &ret, &indexCompare)
	return ret
}

func rightSideView(root *TreeNode, index, level int, ret, indexCompare *[]int) {
	if len(*indexCompare) < level {
		*indexCompare = append(*indexCompare, index)
		*ret = append(*ret, root.Val)
	}else{
		if index > (*indexCompare)[level-1] {
			(*indexCompare)[level-1] = index
			(*ret)[level-1] = root.Val
		}
	}
	if root.Left != nil {
		rightSideView(root.Left, index*2, level+1, ret, indexCompare)
	}
	if root.Right != nil {
		rightSideView(root.Right, index*2+1, level+1, ret, indexCompare)
	}
}