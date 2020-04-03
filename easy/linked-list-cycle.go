func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		if p2 == p1 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}
