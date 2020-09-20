package main

func main() {

}

/**
1.叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
2.当 root 节点左右孩子都为空时，返回 1
3.当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
4.当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值
*/
//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	if root.Left == nil && root.Right == nil {
		return 1
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
