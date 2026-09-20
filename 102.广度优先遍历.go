package leetcodelearn


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	n := len(queue)
	result := [][]int{}
	for n > 0 {
		level := []int{}
		for i := 0; i < n; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		queue = queue[n:]
		n = len(queue)
	}
	return result
}
