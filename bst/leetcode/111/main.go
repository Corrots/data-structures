package main

func main() {

}

//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 {
		return rightDepth + 1
	}
	if rightDepth == 0 {
		return leftDepth + 1
	}
	return min(leftDepth, rightDepth) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
