package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flattenHelp(root)
}

func flattenHelp(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	left := flattenHelp(root.Left)
	right := flattenHelp(root.Right)

	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}

	root.Right = left
	cur.Right = right
	root.Left = nil

	return root
}
