package balancedbinarytree

import (
	"math"
)

// TreeNode
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	val   int
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := math.Abs(float64(depth(root.Left) - depth(root.Right)))

	if diff > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(depth(node.Left), depth(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
