package main

import "fmt"

func leetcode200(grid [][]byte) int {
	var ret int
	leny := len(grid)
	if leny == 0 {
		return 0
	}
	lenx := len(grid[0])
	for j:=0; j<leny; j++ {
		for i:=0; i<lenx; i++ {
			if grid[j][i] == '1' {
				ret++
				checkIsland(grid, i, j, lenx, leny)
			}
		}
	}
	return ret
}

func checkIsland(grid [][]byte, x, y, lenx, leny int) {
	if grid[y][x] == '1' {
		grid[y][x] = 'x'
		if x - 1 >= 0 {
			checkIsland(grid, x-1, y, lenx, leny)
		}
		if y - 1 >= 0 {
			checkIsland(grid, x, y-1, lenx, leny)
		}
		if x + 1 < lenx {
			checkIsland(grid, x+1, y, lenx, leny)
		}
		if y + 1 < leny {
			checkIsland(grid, x, y+1, lenx, leny)
		}
	}
}

func leetcode202(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true { }
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n = n/10
	}
	return sum
}

func leetcode206(head *Node) *Node {
	if head == nil {
		return nil
	}
	var pre, next *Node
	pre = nil
	next = head.Next
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func leetcode215(nums []int, k int) int {
	lens := len(nums) - 1
	for i := lens >> 1; i >= 0; i-- {
		down(nums, i, lens)
	}
	for j := lens; j >= 1; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		lens--
		down(nums, 0, lens)
	}
	return nums[k-1]
}

func down(nums []int, i, lens int) {
	min := i
	if i<<1+1 <= lens && nums[i<<1+1] < nums[min] { //相等最后一个元素
		min = i<<1 + 1
	}
	if i<<1+2 <= lens && nums[i<<1+2] < nums[min] {
		min = i<<1 + 2
	}
	if i != min {
		nums[i], nums[min] = nums[min], nums[i]
		down(nums, min, lens)
	}
}

func leetcode225() {
	o := StackConstructor()
	o.Push(1)
	o.Pop()
}

type MyStack struct {
	Len			int
	Head		*DoubleLinkedNode
	Tail		*DoubleLinkedNode
}


/** Initialize your data structure here. */
func StackConstructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Len++
	o := &DoubleLinkedNode{x, nil, this.Head}
	if this.Head != nil {
		this.Head.Pre = o
	}
	this.Head = o
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ret := this.Head.Val
	this.Head = this.Head.Next
	if this.Head != nil {
		this.Head.Pre = nil
	}
	this.Len--
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Head.Val
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len == 0
}


func leetcode236(root, p, q *TreeNode) *TreeNode {
	var ret *TreeNode
	lowestCommonAncestor(root, p ,q , &ret)
	return ret
}

func lowestCommonAncestor(root, p, q *TreeNode, ret **TreeNode) (bool, bool){
	var left, right bool
	if root == p {
		left = true
	}
	if root == q {
		right = true
	}
	var tmpR, tmpL bool
	if root.Left != nil {
		tmpL, tmpR = lowestCommonAncestor(root.Left, p, q, ret)
	}
	left = left || tmpL
	right = right || tmpR
	if left && right {
		if *ret == nil {
			*ret = root
		}
		return false, false
	}
	if root.Right != nil {
		tmpL, tmpR = lowestCommonAncestor(root.Right, p, q, ret)
	}
	left = left || tmpL
	right = right || tmpR
	if left && right {
		if *ret == nil {
			*ret = root
		}
		return false, false
	}
	return left, right
}

func leetcode283(nums []int) {
	lastIndex := -1
	for k,v := range nums {
		if v == 0 {
			if lastIndex < 0 {
				lastIndex = k
			}
		}else{
			if lastIndex >= 0 {
				nums[k], nums[lastIndex] = nums[lastIndex], nums[k]
				if nums[lastIndex+1] == 0 {
					lastIndex++
				}else{
					lastIndex = k
				}
			}
		}
	}
}

func leetcode289(board [][]int) {
	var count int
	var health bool
	leny := len(board)
	if leny == 0 {
		return
	}
	lenx := len(board[0])
	for j:=0; j<leny; j++ {
		for i:=0; i<lenx; i++ {
			count = 0
			if board[j][i] != 0 {
				health = true
			}else{
				health = false
			}
			if i-1>=0 && j-1>=0 && (board[j-1][i-1] == 1 || board[j-1][i-1] == -1) {
				count++
			}
			if j-1>=0 && (board[j-1][i] == 1 || board[j-1][i] == -1) {
				count++
			}
			if i+1<lenx && j-1>=0 && (board[j-1][i+1] == 1 || board[j-1][i+1] == -1) {
				count++
			}

			if i-1>=0 && (board[j][i-1] == 1 || board[j][i-1] == -1) {
				count++
			}
			if i+1<lenx && (board[j][i+1] == 1 || board[j][i+1] == -1) {
				count++
			}

			if i-1>=0 && j+1<leny && (board[j+1][i-1] == 1 || board[j+1][i-1] == -1) {
				count++
			}
			if j+1<leny && (board[j+1][i] == 1 || board[j+1][i] == -1) {
				count++
			}
			if i+1<lenx && j+1<leny && (board[j+1][i+1] == 1 || board[j+1][i+1] == -1) {
				count++
			}
			fmt.Println(j,i,count)
			if health {
				if count < 2 {
					board[j][i] = -1
				}else if count == 2 || count == 3 {
					continue
				}else{
					board[j][i] = -1
				}
			}else{
				if count == 3 {
					board[j][i] = -2
				}
			}
		}
	}
	//不使用额外空间 就要再遍历一次
	for j:=0; j<leny; j++ {
		for i := 0; i < lenx; i++ {
			if board[j][i] == -1 {
				board[j][i] = 0
			}
			if board[j][i] == -2 {
				board[j][i] = 1
			}
		}
	}
}