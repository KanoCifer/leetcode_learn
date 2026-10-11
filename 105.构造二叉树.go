package leetcodelearn

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root_val := preorder[0]
	root := &TreeNode{
		Val: root_val,
	}
	root_idx := -1
	for i := range inorder {
		if inorder[i] == root_val {
			root_idx = i
		}
	}

	root.Left = buildTree(preorder[1:1+root_idx], inorder[:root_idx])
	root.Right = buildTree(preorder[1+root_idx:], inorder[root_idx+1:])

	return root
}
