package main

func leetcode515(root *TreeNode) []int {
	var ret []int
	getLevelMax(&ret, root, 0)
	return ret
}

func getLevelMax(ret *[]int, root *TreeNode, level int) {
	if len(*ret) - 1 < level {
		*ret = append(*ret, root.Val)
	}
	if root.Val > (*ret)[level] {
		(*ret)[level] = root.Val
	}
	if root.Left != nil {
		getLevelMax(ret, root.Left, level+1)
	}
	if root.Right != nil {
		getLevelMax(ret, root.Right, level+1)
	}
}

func leetcode543(root *TreeNode) int {
	var ret int
	maxRoute(root, &ret)
	return ret
}

func maxRoute(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := maxRoute(root.Left, max)
	r := maxRoute(root.Right, max)

	if l + r > *max {
		*max = l+r
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

func leetcode560(nums []int, k int) int {
	var ret,tmp int
	sum := make(map[int]int)
	sum[0] = 1
	for _,v := range nums {
		tmp += v
		if _,ok := sum[tmp]; ok {
			sum[tmp]++
		} else {
			sum[tmp] = 1
		}
		if count, ok := sum[k-tmp]; ok {
			ret += count
		}
	}
	return ret
}