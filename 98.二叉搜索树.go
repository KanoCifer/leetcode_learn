package leetcodelearn

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max *int) bool
	var min_limit *int
	var max_limit *int

	dfs = func(node *TreeNode, min, max *int) bool {
    if node == nil {
        return true
    }
    if min != nil && node.Val <= *min {
        return false
    }
    if max != nil && node.Val >= *max {
        return false
    }
    return dfs(node.Left, min, &node.Val) && dfs(node.Right, &node.Val, max)
	}
	return dfs(root, min_limit, max_limit)
}
