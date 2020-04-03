/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result bool = true
func isSameTree(p *TreeNode, q *TreeNode) bool {
    result = true
	convertTreeToArray(p,q)
	return result
}


func convertTreeToArray(l *TreeNode, r *TreeNode) {
	fmt.Println(l, r)
	if l != nil && r != nil {
		// fmt.Println(l.Val, r.Val)
		if l.Val != r.Val {
			result = false
		}
		convertTreeToArray(l.Left, r.Left)
		convertTreeToArray(l.Right, r.Right)
	} else if (l != nil && r == nil) || (l == nil && r != nil) {
		result = false
	}
}
