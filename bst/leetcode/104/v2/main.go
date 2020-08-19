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
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for n > 0 {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			n--
		}
		depth++
	}
	return depth
}
