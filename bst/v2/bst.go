package v2

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

// 向二分搜索树中插入元素
func (b *BST) Add(k int) {
	b.root = b.root.add(k)
}

func (b *BST) Contains(k int) bool {
	return b.root.contains(k)
}

func (b *BST) PreOrder() {
	b.root.preOrder()
}

func (b *BST) String() string {
	return b.root.String(0)
}
