package v1

// 递归实现
func RemoveElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = RemoveElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
