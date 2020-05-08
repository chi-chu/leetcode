package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(parseInt("123"))
	//fmt.Println(getRomaNum(27))
	//fmt.Println(checkstring("aab", "c*a*b"))
	//fmt.Println(threeSum([]int{9,8,7,6,5,4,3,2,1}))
	//fmt.Println(generateParenthesis(3))
	//fmt.Println(zstring("LEETCODEISHIRING",4))
	//fmt.Println(longestValidParentheses("()(()"))
	//fmt.Println(combinationSum([]int{2,3,5}, 8))
	//fmt.Println(leetcode43("9","99"))
	//fmt.Println(leetcode45([]int{2,3,1,1,4}))
	//fmt.Println(leetcode46([]int{1,2,3}))
	//fmt.Println(leetcode42([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	//fmt.Println(leetcode300([]int{4,10,4,3,8,9}))
	//printTreeNode(generateNodeTree([]int{1,3,2,5,3,-1,9}))
	//fmt.Println(searchTargetCount([]int{1,2,2,2,3,3,5,8,8}, 3))
	//fmt.Println(leetcode49([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//fmt.Println(leetcode50(2.00000, 10))
	//n := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12},{13,14,15,16}}
	/*n :=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	for _,v := range n {
		fmt.Println(v)
	}
	fmt.Println("-------")
	leetcode48(n)
	for _,v := range n {
		fmt.Println(v)
	}*/
	//fmt.Println(leetcode98(generateNodeTree([]int{0.-1})))
	//fmt.Println(leetcode46([]int{1,2,3}))
	//printNodeList(leetcode23([]*Node{generateNodeList([]int{1,4,5}), generateNodeList([]int{1,3,4}), generateNodeList([]int{2,6})}))
	//printNodeList(leetcode25(generateNodeList([]int{1,2,3,4,5}), 3))

	fmt.Println(leetcode6("LEETCODEISHIRING", 4))
}

/*
* 给定一个整数数组nums和一个目标值target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
* 假设每种输入只会对应一个答案。但是，不能重复利用这个数组中同样的元素。
*
* 示例:
* 给定 nums = [2, 7, 11, 15], target = 9
* 因为 nums[0] + nums[1] = 2 + 7 = 9
* 所以返回 [0, 1]
 */
func numSum(nums []int, target int) [2]int {
	var ret [2]int
	s := make(map[int]int)
	for k,v := range nums {
		s[v] = k
	}
	for k,v := range nums {
		if index, ok := s[target - v];ok {
			ret = [2]int{k, index}
			return ret
		}
	}
	return [2]int{-1,-1}
}

/*实现函数int  atoi(string & src, int & result),字符串转换为整数   例如：
"+123"->123,
"-123"-> -123,
"abs"->0,
"123nndnd34"->123
如果超出整型范围，返回-1*/
func parseInt(s string) int32 {
	var ret int32
	var flag int32 = 1
	for _, v := range s {
		if v == '+' {
			flag = 1
		}else if v == '-' {
			flag = -1
		}else{
			if v > '0' && v < '9' {
				if ret > math.MaxInt32/10 || ret < -math.MaxInt32/10{
					return -1
				}
				ret = ret*10 + (v-'0')
			}else{
				break
			}
		}
	}
	return ret*flag
}


func checkIsland(imap [][]int) int {
	var ret int
	leny := len(imap)
	if leny<0 {
		return 0
	}
	lenx := len(imap[0])
	for i:=0;i<lenx;i++ {
		for j:=0;j<leny;j++ {
			if imap[i][j] == 1 {
				//checkPoint()
			}
		}
	}
	return ret
}

func checkPoint(imap [][]int, x,y,lenx,leny int) {
	//递归找岛屿数量
}

//十进制整型链表相加
func AddTwoNodeList(nodeOne, nodeTwo *Node) *Node {
	nodeOne, _ = reverseNode(nodeOne)
	nodeTwo, _ = reverseNode(nodeTwo)
	var retNode, tmpNodeOne, tmpNodeTwo, finalNode *Node
	var flag bool
	tmpNodeOne = nodeOne
	tmpNodeTwo = nodeTwo
	for tmpNodeOne != nil || tmpNodeTwo != nil {
		if tmpNodeOne != nil && tmpNodeTwo != nil {
			sum := tmpNodeOne.Value + tmpNodeTwo.Value
			if flag {
				sum++
				flag = false
			}
			if sum >= 10 {
				sum -= 10
				flag = true
			}
			tmpNodeOne.Value = sum
			tmpNodeTwo.Value = sum
			finalNode = tmpNodeOne
			tmpNodeOne = tmpNodeOne.Next
			tmpNodeTwo = tmpNodeTwo.Next
			retNode = nodeOne
		}else if tmpNodeOne != nil {
			retNode = nodeOne
			finalNode = tmpNodeOne
			if flag {
				tmpNodeOne.Value++
				flag = false
			}
			if tmpNodeOne.Value >= 10 {
				tmpNodeOne.Value -= 10
				flag = true
			}
			tmpNodeOne = tmpNodeOne.Next
		}else if tmpNodeTwo != nil{
			retNode = nodeTwo
			finalNode = tmpNodeTwo
			if flag {
				tmpNodeTwo.Value++
				flag = false
			}
			if tmpNodeTwo.Value >= 10 {
				tmpNodeTwo.Value -= 10
				flag = true
			}
			tmpNodeTwo = tmpNodeTwo.Next
		}
	}
	if flag && finalNode != nil {
		finalNode.Next = &Node{1,nil}
	}
	return retNode
}

//有序数组找出某个值出现的次数
func searchTargetCount(nums []int, target int) int {
	leftIndex := searchBorder(nums, target, true)
	rightIndex := searchBorder(nums, target, false)
	return rightIndex - leftIndex + 1
}

func searchBorder(nums []int, target int, findLeft bool) int {
	leftIndex := 0
	rightIndex := len(nums) -1
	middle := 0
	for leftIndex + 1 != rightIndex {
		middle = (rightIndex - leftIndex)/2 + leftIndex
		if nums[middle] > target {
			rightIndex = middle
		} else if nums[middle] < target {
			leftIndex = middle
		} else {
			if findLeft {
				rightIndex = middle
			} else {
				leftIndex = middle
			}
		}
	}
	if findLeft {
		if nums[leftIndex] == target {
			return leftIndex
		}
	}
	return rightIndex
}

//猴子选大王（约瑟夫环）
func MonekeyKing(m,n int) int {
	var ret int
	var list []int
	for i:=0;i<m;i++ {
		list = append(list, 1)
	}

	return ret
}