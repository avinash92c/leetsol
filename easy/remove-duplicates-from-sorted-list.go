/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var start *ListNode = head
    var prev *ListNode = start
    var curr *ListNode
    
    if start!=nil && start.Next!=nil{
        curr = start.Next
    }else{
        return head
    }
    
    for prev.Next!=nil{
            if(curr.Val == prev.Val){
                prev.Next = curr.Next
            }else{
                prev = curr
            }
             curr = curr.Next
        }
        return start    
    }
