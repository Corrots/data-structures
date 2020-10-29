package avl

type AVL struct {
	root *TreeNode
	size int
}

func NewAVL() *AVL {
	return &AVL{}
}

func (avl *AVL) Len() int {
	return avl.size
}

func (avl *AVL) IsEmpty() bool {
	return avl.size == 0
}

func (avl *AVL) Add(k int, v interface{}) {
	avl.root.add(k, v)
}

func (avl *AVL) Contains(k int) bool {
	return avl.root.contains(k)
}

func (avl *AVL) Minimum() *TreeNode {
	return avl.root.minimum()
}

func (avl *AVL) Maximum() *TreeNode {
	return avl.root.maximum()
}

func (avl *AVL) Remove(k int) *TreeNode {
	return avl.root.remove(k)
}
