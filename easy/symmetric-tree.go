/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root!=nil{
        if root.Left!=nil && root.Right!=nil{
            return isMirror(root.Left,root.Right)        
        }else if root.Left!=nil || root.Right!=nil{
            return false
        }else{
            return true
        }
    }
    return true
}

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	if l.Val == r.Val {
		return isMirror(l.Right, r.Left) && isMirror(l.Left, r.Right)
	}
	return false
}
