package findmindepth

import (
	"math"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

func minDepth(n *Node) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return 1
	}

	min := math.MaxInt32

	if n.Left != nil {
		min = compare(minDepth(n.Left), min)
	}

	if n.Right != nil {
		min = compare(minDepth(n.Right), min)
	}

	return min + 1
}

func compare(a, b int) int {
	if a < b {
		return a
	}
	return b
}
