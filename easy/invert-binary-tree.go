/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	walk(root)
	return root
}

func walk(node *TreeNode) {
	if node != nil {
		node.Left, node.Right = node.Right, node.Left
		walk(node.Left)
		walk(node.Right)
	}
}
