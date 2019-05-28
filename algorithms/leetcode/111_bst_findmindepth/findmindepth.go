package findmindepth

import "math"

// TreeNode Binary Tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return 1
	}

	min := math.MaxInt32

	if n.Left != nil {
		min = findMin(minDepth(n.Left), min)
	}

	if n.Right != nil {
		min = findMin(minDepth(n.Right), min)
	}

	return min + 1
}

func findMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
