/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
