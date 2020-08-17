package v1

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
// 返回插入新节点后以node为根的二分搜索树
func (n *Node) add(k int, v interface{}) *Node {
	if n == nil {
		return newNode(k, v)
	}
	// k == n.key 不做任何操作
	if k < n.key {
		n.left = n.left.add(k, v)
	} else if k > n.key {
		n.right = n.right.add(k, v)
	}
	return n
}
