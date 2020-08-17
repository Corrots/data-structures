package bst

type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
}

func newNode(k int, v interface{}) *Node {
	return &Node{key: k, value: v}
}

// 向以node为根的二分搜索树中添加元素e
func (n *Node) add(k int, v interface{}) bool {
	if n.key == k {
		return false
	} else if k < n.key && n.left == nil {
		n.left = newNode(k, v)
		return true
	} else if k > n.key && n.right == nil {
		n.right = newNode(k, v)
		return true
	}

	if k < n.key {
		n.left.add(k, v)
	} else {
		n.right.add(k, v)
	}
	return true
}
