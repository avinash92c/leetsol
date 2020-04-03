/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	currNode := head
	var next, prev *ListNode

	for currNode != nil {
		next = currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = next
	}

	return prev
}
