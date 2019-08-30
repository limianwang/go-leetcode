package rangerum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}

	if root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}
