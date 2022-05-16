package a104

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	res := maxDepth(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		}})

	assert.Equal(t, 3, res)

}
