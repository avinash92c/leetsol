/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		if root.Left == nil && root.Right == nil {
			return 1
		}
        // fmt.Println(subdepth(root.Left, root.Right))
		return subdepth(root.Left, root.Right) + 1
	}
}

func subdepth(l *TreeNode, r *TreeNode) int {
	if l == nil && r == nil {
		return 0
	} else if l == nil || r == nil {
		if l != nil {
			return subdepth(l.Left, l.Right) + 1
		}
		if r != nil {
			return subdepth(r.Left, r.Right) + 1
		}
	} //else {
		return max(subdepth(l.Left, l.Right), subdepth(r.Left, r.Right)) + 1
	//}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
