func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1, curr2 := l1, l2
	// var curr *ListNode
	node := &ListNode{}
	curr := node
	sum := 0

	for curr1 != nil || curr2 != nil {
		if curr1 != nil {
			sum += curr1.Val
			curr1 = curr1.Next
		}
		if curr2 != nil {
			sum += curr2.Val
			curr2 = curr2.Next
		}

		fmt.Println(sum)

		v := &ListNode{Val: sum % 10}
		sum /= 10
		curr.Next = v
		curr = curr.Next
	}

	if sum != 0 {
		curr.Next = &ListNode{Val: sum % 10}
	}

	return node.Next
}
