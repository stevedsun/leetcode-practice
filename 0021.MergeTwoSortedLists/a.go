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

	if p1.Val < p2.Val {
		p1.Next = mergeTwoLists(p1.Next, p2)
		return p1
	} else {
		p2.Next = mergeTwoLists(p1, p2.Next)
		return p2
	}
}
