package a21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2

	if p1 == nil {
		return p2
	}

	if p2 == nil {
		return p1
	}

	var result *ListNode
	var resHead *ListNode

	if p1.Val < p2.Val {
		result = p1
		resHead = result
		p1 = p1.Next
	} else {
		result = p2
		resHead = result
		p2 = p2.Next
	}

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			result.Next = p1
			result = result.Next
			p1 = p1.Next
		} else {
			result.Next = p2
			result = result.Next
			p2 = p2.Next
		}
	}

	if p1 == nil && p2 != nil {
		result.Next = p2
	}

	if p2 == nil && p1 != nil {
		result.Next = p1
	}

	return resHead
}
