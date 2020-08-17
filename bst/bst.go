package bst

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return &BST{root: nil, size: 0}
}

func (b *BST) Len() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(k int, v interface{}) {
	if b.root == nil {
		b.root = newNode(k, v)
	} else {
		if ok := b.root.add(k, v); ok {
			b.size++
		}
	}
}
