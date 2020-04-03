/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	walksum(root, &sum)
	return sum
}

func walksum(node *TreeNode, sum *int) {
	if node != nil {
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			*sum += node.Left.Val
		}
		walksum(node.Left, sum)
		walksum(node.Right, sum)
	}
}
