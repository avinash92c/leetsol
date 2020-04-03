/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	rootNode := head
	currNode := head
	var prevNode *ListNode
	for currNode != nil {
		if currNode.Val == val {
			if prevNode != nil {
				prevNode.Next = currNode.Next
			} else {
				rootNode = currNode.Next
			}
		} else {
			prevNode = currNode
		}
		currNode = currNode.Next
	}
	return rootNode
}
