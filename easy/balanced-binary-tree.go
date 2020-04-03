/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//CHECK MAX DEPTH OF LEFT TREE
	ld := getSubtreeHeight(root.Left)
	//CHECK MAX DEPTH OF RIGHT TREE
	rd := getSubtreeHeight(root.Right)
	//COMPARE
	// fmt.Println(ld, rd)
	return math.Abs(float64(ld-rd)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getSubtreeHeight(tree *TreeNode) int{
	if tree == nil {
		return 0
	}
	return max(getSubtreeHeight(tree.Left), getSubtreeHeight(tree.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
