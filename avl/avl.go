package avl

type AVL struct {
	root *TreeNode
	size int
}

func NewAVL() *AVL {
	return &AVL{}
}

// 判断该二叉树是否是一棵二叉搜索树(通过中序遍历判断)
func (avl *AVL) IsBST() bool {
	var keys []int
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.left)
		keys = append(keys, node.key)
		inOrder(node.right)
	}
	inOrder(avl.root)
	for i := 1; i < len(keys); i++ {
		if keys[i-1] > keys[i] {
			return false
		}
	}
	return true
}

// 判断该二叉树是否为平衡二叉树
func (avl *AVL) IsBalanced() bool {
	return avl.root.isBalanced()
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
