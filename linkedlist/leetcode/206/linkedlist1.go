package leetcode

// 反转链表-递归实现
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
}
