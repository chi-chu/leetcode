package main

import "fmt"

type Node struct {
	Value	int
	Next	*Node
}

type TreeNode struct {
	Val 	int
	Left	*TreeNode
	Right	*TreeNode
}

//生成单链表
func generateNodeList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	ret := &Node{nums[0],nil}
	tmp := ret
	for k,v := range nums {
		if k == 0 {
			continue
		}
		tmp.Next = &Node{v, nil}
		tmp = tmp.Next
	}
	return ret
}

//null的点 值为 -1
func generateNodeTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	ret := &TreeNode{nums[0], nil, nil}
	tmp := []*TreeNode{ret}
	for k,v := range nums {
		if k == 0 {
			continue
		}
		if v != -1 {
			index := (k-1)/2
			current := &TreeNode{v, nil, nil}
			if tmp[index] != nil {
				if k % 2 == 0 {
					tmp[index].Right = current
				}else{
					tmp[index].Left = current
				}
			}
			tmp = append(tmp, current)
		}else{
			tmp = append(tmp, nil)
		}
	}
	return ret
}


//打印单链表
func printNodeList(n *Node) {
	var show []int
	for n != nil {
		show = append(show, n.Value)
		n = n.Next
	}
	fmt.Println(show)
}

//链表反转
func reverseNode(root *Node) (*Node, int) {
	var lens int
	var next, pre *Node
	pre = nil
	for root != nil {
		lens++
		next = root.Next
		root.Next = pre
		pre = root
		root = next
	}
	return pre, lens
}

func printTreeNode(n *TreeNode) {
	if n == nil || n.Val == -1 {
		fmt.Println("TreeNode is nil ")
		return
	}
	var view [][]int
	parseTree(&view, n, 0)
	for _,v := range view {
		fmt.Println(v)
	}
}

func parseTree(view *[][]int, root *TreeNode, level int) {
	if len(*view) - 1 < level {
		*view = append(*view, []int{})
	}
	(*view)[level] = append((*view)[level], root.Val)
	if root.Left != nil {
		parseTree(view, root.Left, level+1)
	}
	if root.Left == nil && root.Right != nil {
		if len(*view) - 2 < level {
			*view = append(*view, []int{})
		}
		(*view)[level+1] = append((*view)[level+1], -1)
	}
	if root.Right != nil {
		parseTree(view, root.Right, level+1)
	}
}

//快速排序
func quickSort(n []int, left, right int) {
	if left < right {
		index := getIndex(n, left, right)
		quickSort(n, left, index-1)
		quickSort(n, index+1, right)
	}
}

func getIndex(n []int, left, right int) int {
	tmp := n[left]
	for left < right {
		for n[left] < tmp {
			left += 1
		}
		for n[right] > tmp {
			right -= 1
		}
		swtmp := n[left]
		n[left] = n[right]
		n[right] = swtmp
	}
	n[left] = tmp
	return left
}
//快速排序end

//二叉树反转
func reverseTreeNode(root *TreeNode) {
	if root.Right == nil && root.Left == nil {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		reverseTreeNode(root.Left)
	}
	if root.Right != nil {
		reverseTreeNode(root.Right)
	}
}