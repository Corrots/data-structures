package solution

import (
	"fmt"
	"log"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 以arr为参数，创建一个链表，当前的ListNode为链表头节点
func NewListNode(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		log.Fatal("arr can not be empty")
	}
	listNode := &ListNode{Val: arr[0]}
	cur := listNode
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return listNode
}

// 以当前节点为头节点的链表信息字符串
func (l *ListNode) String() string {
	var res strings.Builder
	cur := l
	for cur != nil {
		res.WriteString(fmt.Sprintf("%d->", cur.Val))
		cur = cur.Next
	}
	res.WriteString("nil")
	return res.String()
}
