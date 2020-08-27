package main

func main() {

}

//https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	delNode := node.Next
	node.Next = delNode.Next
	delNode = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
