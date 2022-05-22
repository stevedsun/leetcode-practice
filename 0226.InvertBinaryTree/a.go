package a226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	leftTree := invertTree(root.Left)
	rightTree := invertTree(root.Right)

	root.Left = rightTree
	root.Right = leftTree

	return root
}
