package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = []string{}

func binaryTreePaths(root *TreeNode) []string {
	result = []string{}
	dfs(root, "")
	return result
}

func dfs(cur *TreeNode, path string) {

	if cur.Left == nil && cur.Right == nil {
		if path == "" {
			result = append(result, strconv.Itoa(cur.Val))
		} else {
			result = append(result, path+"->"+strconv.Itoa(cur.Val))
		}

		return
	}

	if cur.Left != nil {
		if path == "" {
			dfs(cur.Left, strconv.Itoa(cur.Val))
		} else {
			dfs(cur.Left, path+"->"+strconv.Itoa(cur.Val))
		}

	}

	if cur.Right != nil {
		if path == "" {
			dfs(cur.Right, strconv.Itoa(cur.Val))
		} else {
			dfs(cur.Right, path+"->"+strconv.Itoa(cur.Val))
		}

	}

}
