package set

import v2 "github.com/corrots/data-structures/bst/v2"

type BSTSet struct {
	BST *v2.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{BST: v2.NewBST()}
}

func (s *BSTSet) Add(e interface{}) {
	s.BST.Add(e.(int))
}

func (s *BSTSet) Remove(e interface{}) {
	s.BST.Remove(e.(int))
}

func (s *BSTSet) Contains(e interface{}) bool {
	return s.BST.Contains(e.(int))
}

func (s *BSTSet) Len() int {
	return s.BST.Len()
}

func (s *BSTSet) IsEmpty() bool {
	return s.BST.IsEmpty()
}
