package main

import "fmt"

//================================== 单链表
type Node struct {
	Value	int
	Next	*Node
}

//生成单链表
func NewNodeList(nums []int) *Node {
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

func (n *Node) Len() int {
	ret := 0
	for n != nil {
		n = n.Next
		ret++
	}
	return ret
}

func (n *Node) RemoveByNode(p *Node) {
	if n == nil || p == nil {
		return
	}
	if n == p {
		*n = *n.Next
		return
	}
	var pre *Node
	tmp := n
	for tmp.Next != p {
		pre = tmp
		tmp = tmp.Next
	}
	if pre != nil {
		pre.Next = tmp.Next.Next
	}
}

func (n *Node) RemoveByVal(v int) {
	if n == nil {
		return
	}
	var pre *Node
	tmp := n
	for tmp.Next != nil && tmp.Next.Value != v {
		pre = tmp
		tmp = tmp.Next
	}
	if pre != nil && tmp.Next != nil {
		pre.Next = tmp.Next.Next
	}
}

func (n *Node) AddNodeHead(v int) {
	*n = Node{v,n}
}

func (n *Node) AddNodeTail(v int) {
	if n == nil {
		*n = Node{v,nil}
		return
	}
	tmp := n
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{v,nil}
}

//链表反转
func (n *Node) Reverse() {
	var next, pre *Node
	root := n
	pre = nil
	for root != nil {
		next = root.Next
		root.Next = pre
		pre = root
		root = next
	}
	if pre != nil {
		*n = *pre
	}
}

//打印单链表
func (n *Node) Print() {
	var show []int
	for n != nil {
		show = append(show, n.Value)
		n = n.Next
	}
	fmt.Println(show)
}

//====================================== 双向链表
type DoubleLinkedNode struct {
	Val		int
	Pre		*DoubleLinkedNode
	Next	*DoubleLinkedNode
}

func NewDoubleLinkedList(nums []int) (*DoubleLinkedNode, *DoubleLinkedNode) {
	lens := len(nums)
	if lens == 0 {
		return nil, nil
	}
	ret := &DoubleLinkedNode{nums[0], nil, nil}
	if lens == 1 {
		return ret, ret
	}
	var pre, tmp *DoubleLinkedNode
	tmp = ret
	pre = nil
	for k,v := range nums {
		if k == 0 {
			continue
		}else{
			tmp.Pre = pre
			tmp.Next = &DoubleLinkedNode{v, nil, nil}
			pre = tmp
			tmp = tmp.Next
		}
	}
	return ret, pre
}

func (d *DoubleLinkedNode) Len() int {
	var ret int
	for d != nil {
		d = d.Next
		ret ++
	}
	return ret
}

func (d *DoubleLinkedNode) AddHead(v int) {
	*d = DoubleLinkedNode{v, nil, d}
}

func (d *DoubleLinkedNode) AddTail(v int) {
	if d == nil {
		*d = DoubleLinkedNode{v, nil, nil}
		return
	}
	tmp := d
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &DoubleLinkedNode{v, tmp, nil}
}

func (d *DoubleLinkedNode) Remove(v int, all bool) {
	tmp := d
	for {
		if tmp == nil {
			return
		}
		if tmp.Val == v {
			tmp.Pre.Next = tmp.Next
			if tmp.Next != nil {
				tmp.Next.Pre = tmp.Pre
			}
			if all {
				tmp = tmp.Next
			}else{
				return
			}
		}else{
			tmp = tmp.Next
		}
	}
}

//====================================== 树
type TreeNode struct {
	Val 	int
	Left	*TreeNode
	Right	*TreeNode
}

//TODO remove treeNode ~~~
/*func (t *TreeNode) remove(v int, AVL bool) {
	if AVL {
		tmp := t
		for tmp != nil && tmp.Val != v {
			if tmp.Val < v {
				tmp = tmp.Right
			}else{
				tmp = tmp.Left
			}
		}
		if tmp != nil {

		}
	}else{

	}
}*/

//null的点 值为 -1
func NewNodeTree(nums []int) *TreeNode {
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

//二叉树反转
func (t *TreeNode) Reverse() {
	if t.Right == nil && t.Left == nil {
		return
	}
	tmp := t.Left
	t.Left = t.Right
	t.Right = tmp
	if t.Left != nil {
		t.Left.Reverse()
	}
	if t.Right != nil {
		t.Right.Reverse()
	}
}

func (t *TreeNode) PrintTreeNode() {
	if t == nil || t.Val == -1 {
		fmt.Println("TreeNode is nil ")
		return
	}
	var view [][]int
	parseTree(&view, t, 0)
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

//====================================== 其他
//快速排序
func QuickSort(n []int, left, right int) {
	if left < right {
		index := getIndex(n, left, right)
		QuickSort(n, left, index-1)
		QuickSort(n, index+1, right)
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