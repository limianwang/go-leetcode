package inorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := make([]int, 0)
	inorder(root, &vals)

	tree := &TreeNode{
		Val: 0,
	}

	curr := tree

	for _, v := range vals {
		curr.Right = &TreeNode{
			Val: v,
		}

		curr = curr.Right
	}

	return tree.Right
}

func inorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}
