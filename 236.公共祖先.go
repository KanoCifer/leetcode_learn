package leetcodelearn

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(root *TreeNode) (bool, bool, *TreeNode)

	dfs = func(root *TreeNode) (bool, bool, *TreeNode){
		// 边界判断
		if root == nil {
			return false, false, nil
		}

		left_p, left_q, left_x := dfs(root.Left)
		right_p, right_q, right_x := dfs(root.Right)

		has_p := left_p || right_p || root.Val == p.Val
		has_q := left_q || right_q || root.Val == q.Val

		var x *TreeNode
		if left_x != nil {
			x = left_x
		} else if right_x != nil {
			x = right_x
		} else if has_p && has_q {
			x = root
		}


		return has_p, has_q, x
	}

	_, _, x := dfs(root)
	return x
}
