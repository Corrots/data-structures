package main

import "fmt"

func main() {
	//nums := []int{1, 0, 1}
	nums := []int{1, 2, 2, 1}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	fmt.Println(isPalindrome(l1))
}

//https://leetcode-cn.com/problems/palindrome-linked-list-lcci/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	leftTail := getLeftTail(head)
	//PrintLinkedList(leftTail)
	reverse := reverseList(leftTail.Next)
	//PrintLinkedList(reverse)
	p := head
	q := reverse
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

// 获取以head为头的链表前半部分的尾节点
// 快慢指针
func getLeftTail(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 反转以head为头的链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func CreateLinkedList(nums []int, n int) *ListNode {
	if n == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = NewListNode(nums[i])
		cur = cur.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("nil\n")
}
