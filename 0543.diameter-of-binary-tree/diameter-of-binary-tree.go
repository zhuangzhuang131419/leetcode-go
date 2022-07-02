package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	height(root.Right)
	height(root.Left)

	return max(max(diameterOfBinaryTree(root.Right), diameterOfBinaryTree(root.Left)), height(root.Left)+height(root.Right))

}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return max(height(root.Left), height(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
