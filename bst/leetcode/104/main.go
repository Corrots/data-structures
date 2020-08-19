package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func maxDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
