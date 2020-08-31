package main

func main() {

}

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*Command
	stack = append(stack, &Command{isPrint: false, node: root})
	for len(stack) > 0 {
		sz := len(stack)
		cmd := stack[sz-1]
		stack = stack[:sz-1]
		if cmd.isPrint {
			res = append(res, cmd.node.Val)
		} else {
			if cmd.node.Right != nil {
				stack = append(stack, &Command{isPrint: false, node: cmd.node.Right})
			}
			if cmd.node.Left != nil {
				stack = append(stack, &Command{isPrint: false, node: cmd.node.Left})
			}
			stack = append(stack, &Command{isPrint: true, node: cmd.node})
		}
	}
	return res
}

type Command struct {
	isPrint bool
	node    *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
