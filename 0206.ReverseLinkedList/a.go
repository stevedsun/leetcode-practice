package a206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p1 *ListNode = nil
	var p3 *ListNode = nil
	p2 := head

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	return p1
}
