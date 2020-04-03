/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head==nil||head.Next == nil {
		return true
	}
	valarr := make([]int, 0)
	curr := head
	for curr != nil {
		valarr = append(valarr, curr.Val)
		curr = curr.Next
	}
    fmt.Println(valarr)
	i, j := 0, len(valarr)-1
	for i < j {
        fmt.Println(valarr[i],valarr[j])
		if valarr[i] != valarr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
