/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    //generate array from list
    nums := make([]int,0)
    curr:= head
    for curr!=nil{
        nums = append(nums,curr.Val)
        curr = curr.Next
    }
    //run through array and generate tree
    return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//take mid and place at start node
	idx := len(nums) / 2
	return &TreeNode{
		Val:   nums[idx],
		Left:  sortedArrayToBST(nums[:idx]),
        Right: sortedArrayToBST(nums[idx+1:]),
	}
}
