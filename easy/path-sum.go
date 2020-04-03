/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
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
var resultSum = 0
func hasPathSum(root *TreeNode, sum int) bool {
    resultSum=0
	return inOrderTraversal(root, sum)
}
func inOrderTraversal(node *TreeNode, sum int) bool {
	result := false
	if node == nil {
		return result
	}
    resultSum += node.Val
     fmt.Println("Node Val")
     fmt.Println(node.Val)
	fmt.Println("Sum")
	fmt.Println(resultSum)
	if node.Left != nil {
		fmt.Println("Calling Left")
		result = inOrderTraversal(node.Left, sum)
		if result {
			return result
		}
	}
	if node.Right != nil {
	fmt.Println("Calling Right")
		result = inOrderTraversal(node.Right, sum)
		if result {
			return result
		}
	}
	if resultSum != sum {
        fmt.Println("Removing Current Val-")
        fmt.Println(resultSum-node.Val)
        resultSum -= node.Val
	} else {
		if resultSum == sum && node.Left == nil && node.Right == nil {
            fmt.Println("Returning True")
			return true
		} else {
			resultSum -= node.Val
		 }
	}
	return result
}
