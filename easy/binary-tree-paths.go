/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	walkpath(root, "", &result)
	return result
}

func walkpath(node *TreeNode, currpath string, result *[]string) {
	if node != nil {
		if len(currpath) == 0 {
			currpath = strconv.Itoa(node.Val)
		} else {
			currpath = currpath + "->" + strconv.Itoa(node.Val)
		}

		if node.Left == nil && node.Right == nil {
			*result = append(*result, currpath)
		} else {
			walkpath(node.Left, currpath, result)
			walkpath(node.Right, currpath, result)
		}
	}
}
