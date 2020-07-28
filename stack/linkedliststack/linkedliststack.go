package linkedliststack

import (
	"strings"

	"github.com/corrots/data-structures/stack"
)

type LinkedListStack struct {
	list *LinkedList
}

func NewStack() stack.Stack {
	return &LinkedListStack{list: NewLinkedList()}
}

func (l *LinkedListStack) Push(e interface{}) {
	l.list.AddFirst(e)
}

func (l *LinkedListStack) Pop() interface{} {
	return l.list.RemoveFirst()
}

func (l *LinkedListStack) Peek() interface{} {
	return l.list.GetFirst()
}

func (l *LinkedListStack) Len() int {
	return l.list.Len()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.list.Len() == 0
}

func (l *LinkedListStack) String() string {
	var res strings.Builder
	res.WriteString("Stack: top ")
	res.WriteString(l.list.String())
	return res.String()
}
