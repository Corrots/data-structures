package set

import (
	"github.com/corrots/data-structures/linkedlist/linkedlist"
)

type LinkedListSet struct {
	list *linkedlist.LinkedList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{
		list: linkedlist.New(),
	}
}

// 基于链表实现的Set，需先查重
func (s *LinkedListSet) Add(i interface{}) {
	if !s.Contains(i) {
		s.list.AddFirst(i)
	}
}

func (s *LinkedListSet) Remove(i interface{}) {
	s.list.Remove(i.(int))
}

func (s *LinkedListSet) Contains(i interface{}) bool {
	return s.list.Contains(i)
}

func (s *LinkedListSet) Len() int {
	return s.list.Len()
}

func (s *LinkedListSet) IsEmpty() bool {
	return s.list.IsEmpty()
}
