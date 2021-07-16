package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 后序
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

// 中序
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
