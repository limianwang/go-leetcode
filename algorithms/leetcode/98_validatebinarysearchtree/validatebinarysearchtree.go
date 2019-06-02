package validatebinarysearchtree

import "math"

// TreeNode is a node in a Tree
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

/**
I know that
- left is always smaller than node
- right is always greater than node

I will need to go through all nodes in order to determine if this is a
BST. => Recursion

I know that I want to check my left then also my right.

Is a nil node still considered a valid tree?
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	minLeft := math.MinInt16
	maxRight := math.MaxInt16

	return check(root.Left, minLeft, root.Val) && check(root.Right, root.Val, maxRight)
}

func check(node *TreeNode, left int, right int) bool {
	if node == nil {
		return true
	}

	if left < node.Val && node.Val < right {
		return check(node.Left, left, node.Val) && check(node.Right, node.Val, right)
	}

	return false
}
