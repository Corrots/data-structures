package bstmap

type TreeNode struct {
	key   int
	val   interface{}
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(k int, v interface{}) *TreeNode {
	return &TreeNode{key: k, val: v}
}

// 根据k来获取指定的节点
func (n *TreeNode) getTreeNode(k int) *TreeNode {
	if n == nil {
		return nil
	}
	if k == n.key {
		return n
	}
	if k < n.key {
		return n.left.getTreeNode(k)
	} else {
		return n.right.getTreeNode(k)
	}
}

// 向以node为根的二分搜索树中添加元素(k, v)
// 返回插入新节点后以node为根的二分搜索树
func (n *TreeNode) add(k int, v interface{}) *TreeNode {
	if n == nil {
		return newTreeNode(k, v)
	}
	// k == n.key时更新操作
	if k == n.key {
		n.val = v
	} else if k < n.key {
		n.left = n.left.add(k, v)
	} else {
		n.right = n.right.add(k, v)
	}
	return n
}

// 查找以node为根的二分搜索树中是否包含节点k
//func (n *TreeNode) contains(k int) bool {
//	if n == nil {
//		return false
//	}
//	if k == n.key {
//		return true
//	} else if k < n.key {
//		return n.left.contains(k)
//	} else {
//		return n.right.contains(k)
//	}
//}

// 获取以n为根节点的二分搜索树中的最小值
func (n *TreeNode) minimum() *TreeNode {
	if n.left == nil {
		return n
	}
	return n.left.minimum()
}

// 删除以n为根节点的二分搜索树中的最小值
func (n *TreeNode) removeMin() *TreeNode {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		return rightNode
	}
	n.left = n.left.removeMin()
	return n
}

// 在以n为根的二分搜索树中删除节点k
func (n *TreeNode) remove(k int) *TreeNode {
	if n == nil {
		return nil
	}
	if k < n.key { // 在n的左子树中查找删除k
		n.left = n.left.remove(k)
		return n
	} else if k > n.key { // 在n的右子树中查找删除k
		n.right = n.right.remove(k)
		return n
	} else { // k == n.key
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			return rightNode
		}
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			return leftNode
		}
		// 左右子树 != nil
		// 1.找到n的后继节点;
		// 2.将删除后继节点后的n的右子树作为后继节点的右子树;
		// 3.后继节点的左子树 = n的左子树
		successor := n.right.minimum()
		successor.right = n.right.removeMin()
		successor.left = n.left
		n.left, n.right = nil, nil
		return successor
	}
}
