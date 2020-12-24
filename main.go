package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(leetcode102(NewNodeTree([]int{1,7,9,-1,-1,1,7,-1,-1,-1,-1,-1,8,9})))
	//fmt.Println(leetcode199(NewNodeTree([]int{1,2,3,-1,5,-1,4})))
	//fmt.Println(leetcode200([][]byte{{'1','1','0','0','0'}, {'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}))
	//n := [][]int{{0,1,0}, {0,0,1},{1,1,1}, {0,0,0}}
	//leetcode289(n)
	//fmt.Println(n)
	//fmt.Println(leetcode322([]int{186,419,83,408},6249))
	//fmt.Println(leetcode72("horse","ros"))
	//fmt.Println(leetcode560([]int{1,1,1},2))
	//fmt.Println(MonekeyKing(10, 3))
	//leetcode25(NewNodeList([]int{1,2}), 1).Print()
	//fmt.Println(leetcode41([]int{3,4,-1,1}))
	//n := []int{1,2,1,2,3,7}
	//fmt.Println(leetcode914([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,5,5,5,5,5,5,5,5,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,7,7,7,7,7,7,7,7,7,7,7,7,8,8,8,8,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,11,11,11,11,12,12,12,12,12,12,12,12,12,12,12,12,12,12,12,12,13,13,13,13,13,13,13,13,14,14,14,14,14,14,14,14,15,15,15,15,15,15,15,15,15,15,15,15,16,16,16,16,17,17,17,17,18,18,18,18,19,19,19,19}))
	//fmt.Println(leetcode1103(10,3))
	//fmt.Println(leetcode945([]int{2,2,2,1}))
	//fmt.Println(leetcode994([][]int{{0}}))
	//fmt.Println(leetcode1013([]int{29,31,27,-10,-67,22,15,-1,-16,-29,59,-18,48}))
	//fmt.Println(leetcode5("babad"))
	//fmt.Println(leetcode394( "2[abc]3[cd]ef"))
	//fmt.Println(leetcode93("25525511135"))
	fmt.Println(leetcode84([]int{2,1,5,6,2,3}))
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

//十进制整型链表相加
func AddTwoNodeList(nodeOne, nodeTwo *Node) *Node {
	nodeOne.Reverse()
	nodeTwo.Reverse()
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
func MonekeyKing(m, n int) int {
	if m < 1 {
		return -1
	}
	var pre, index, count int
	var list []int
	for i:=0; i<m; i++ {
		list = append(list, i+1)
	}
	list[m-1] = 0
	pre = m - 1
	for m != 1{
		if list[index] != -1 {
			count++
			if count == n {
				list[pre] = list[index]
				list[index] = -1
				index = list[pre]
				count = 0
				m--
			} else {
				pre = index
				index = list[index]
			}
		}
		fmt.Println(list)
	}
	return index+1
}