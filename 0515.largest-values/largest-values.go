package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodes := []*TreeNode{}

	nodes = append(nodes, root)

	result := []int{}

	for len(nodes) != 0 {

		maxValue := 0
		size := len(nodes)
		for i := 0; i < size; i++ {
			treeNode := nodes[i]
			if treeNode == nil {
				continue
			}
			if treeNode.Val > maxValue {
				maxValue = treeNode.Val
			}
			if treeNode.Left != nil {
				nodes = append(nodes, treeNode.Left)
			}

			if treeNode.Right != nil {
				nodes = append(nodes, treeNode.Right)
			}
		}
		nodes = nodes[size:]
		result = append(result, maxValue)
	}

	return result

}
