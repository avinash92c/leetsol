func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	currnode := head
	data := make([]*ListNode, 0)
	idx := 0
	for currnode != nil {
		data = append(data,currnode)
		currnode = currnode.Next
		idx++
	}

	k = k % idx
	if k == 0 {
		return data[0]
	}
    fmt.Println(k)
    
	data[idx-1].Next = data[0]
    
	data[idx-k-1].Next = nil

	return data[idx-k]
}
