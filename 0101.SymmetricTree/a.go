package a101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	left := frontSearch(root.Left)
	right := backSearch(root.Right)

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func frontSearch(root *TreeNode) []int {
	if root == nil {
		return []int{-101}
	}

	left := frontSearch(root.Left)
	right := frontSearch(root.Right)

	result := []int{}

	result = append(result, left...)
	result = append(result, right...)
	result = append(result, root.Val)

	return result
}

func backSearch(root *TreeNode) []int {
	if root == nil {
		return []int{-101}
	}

	right := backSearch(root.Right)
	left := backSearch(root.Left)

	result := []int{}

	result = append(result, right...)
	result = append(result, left...)
	result = append(result, root.Val)

	return result
}
