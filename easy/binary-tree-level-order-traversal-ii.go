func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	traversetree(root, 0, &result)
	generateresult(result)
	return result
}

func traversetree(node *TreeNode, depth int, result *[][]int) {
	if node == nil {
		return
	}

	if len(*result) <= depth {
		*result = append(*result, []int{})
	}

	(*result)[depth] = append((*result)[depth], node.Val)
	traversetree(node.Left, depth+1, result)
	traversetree(node.Right, depth+1, result)
}

func generateresult(result [][]int) {
    // fmt.Println(result)
	for i := 0; i < len(result)/2; i++ {
        // fmt.Println(result[i],result[len(result)-1-i])
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
}
