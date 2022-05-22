package a234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slowPtr := head
	fastPtr := head

	for fastPtr != nil {
		slowPtr = slowPtr.Next
		if fastPtr.Next == nil {
			break
		}
		fastPtr = fastPtr.Next.Next
	}

	var p1, p2, p3 *ListNode
	p1 = nil
	p2 = slowPtr

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	h := head

	for i := p1; i != nil; i = i.Next {
		if i.Val != h.Val {
			return false
		}

		h = h.Next
	}

	return true
}
