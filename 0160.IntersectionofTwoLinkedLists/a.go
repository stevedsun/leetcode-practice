package a160

type ListNode struct {
	Val  int
	Next *ListNode
}

// Not best solution

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)

	if headA == nil || headB == nil {
		return nil
	}

	for i := headA; i != nil; i = i.Next {
		m[i] = 1
	}

	for i := headB; i != nil; i = i.Next {
		if _, ok := m[i]; ok {
			return i
		} else {
			m[i] = 1
		}
	}

	return nil
}
