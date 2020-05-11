package main

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