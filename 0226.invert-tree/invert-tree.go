package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	left := root.Left
	right := root.Right

	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}
