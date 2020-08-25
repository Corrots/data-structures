package v2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表-递归实现
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
}
