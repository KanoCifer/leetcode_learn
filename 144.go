package leetcodelearn

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	return result
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		result = append(result, node.Val)
		fn(node.Right)
	}
	fn(root)
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		fn(node.Right)
		result = append(result, node.Val)
	}
	fn(root)
	return result
}
