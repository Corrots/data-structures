package v2

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	key   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(k int) *TreeNode {
	return &TreeNode{key: k}
}

// 向以node为根的二分搜索树中添加元素e
// 返回插入新节点后以node为根的二分搜索树
func (n *TreeNode) add(k int) *TreeNode {
	if n == nil {
		return newTreeNode(k)
	}
	// k == n.key 不做任何操作
	if k < n.key {
		n.left = n.left.add(k)
	} else if k > n.key {
		n.right = n.right.add(k)
	}
	return n
}

// 查找以node为根的二分搜索树中是否包含节点k
func (n *TreeNode) contains(k int) bool {
	if n == nil {
		return false
	}
	if k == n.key {
		return true
	} else if k < n.key {
		return n.left.contains(k)
	} else {
		return n.right.contains(k)
	}
}

// 以node为根节点的二分搜索树的前序遍历
func (n *TreeNode) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.key)
	n.left.preOrder()
	n.right.preOrder()
}

// 以node为根节点的中序遍历
func (n *TreeNode) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.key)
	n.right.inOrder()
}

func (n *TreeNode) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.key)
}

// 获取以n为根节点的二分搜索树中的最小值
func (n *TreeNode) minimum() *TreeNode {
	if n.left == nil {
		return n
	}
	return n.left.minimum()
}

// 获取以n为根节点的二分搜索树中的最大值
func (n *TreeNode) maximum() int {
	if n.right == nil {
		return n.key
	}
	return n.right.maximum()
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

// 删除以n为根节点的二分搜索树中的最小值
func (n *TreeNode) removeMax() *TreeNode {
	if n.right == nil {
		nodeLeft := n.left
		n.left = nil
		return nodeLeft
	}
	n.right = n.right.removeMax()
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

func (n *TreeNode) String(depth int) string {
	var res strings.Builder
	if n == nil {
		res.WriteString(depthString(depth) + "nil\n")
		return res.String()
	}
	res.WriteString(fmt.Sprintf("%s%d\n", depthString(depth), n.key))
	res.WriteString(n.left.String(depth + 1))
	res.WriteString(n.right.String(depth + 1))
	return res.String()
}

func (n *TreeNode) generateBSTString(depth int, res *strings.Builder) {
	if n == nil {
		res.WriteString(depthString(depth) + "nil\n")
		return
	}
	res.WriteString(fmt.Sprintf("%s%d\n", depthString(depth), n.key))
	n.left.generateBSTString(depth+1, res)
	n.right.generateBSTString(depth+1, res)
}

func depthString(depth int) string {
	var res strings.Builder
	res.WriteString(strings.Repeat("--", depth))
	return res.String()
}
