package a21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	res := mergeTwoLists(l1, l2)

	for _, i := range []int{1, 1, 2, 3, 4, 4} {
		assert.Equal(t, i, res.Val)
		res = res.Next
	}
}
