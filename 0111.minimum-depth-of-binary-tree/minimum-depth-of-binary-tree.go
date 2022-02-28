package _111_minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return BFS(root)
}

func BFS(start *TreeNode) int {
	var q []*TreeNode
	visited := map[*TreeNode]bool{}

	step := 1

	q = append(q, start)
	visited[start] = true

	for len(q) > 0 {

		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if cur.Left == nil && cur.Right == nil {
				return step
			}

			if _, exist := visited[cur.Left]; !exist {
				if cur.Left != nil {
					q = append(q, cur.Left)
					visited[cur.Left] = true
				}
			}

			if _, exist := visited[cur.Right]; !exist {
				if cur.Right != nil {
					q = append(q, cur.Right)
					visited[cur.Right] = true
				}
			}
		}
		step++

	}

	return step

}
