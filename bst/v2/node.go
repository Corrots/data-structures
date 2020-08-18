package v2

import (
	"fmt"
	"strings"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

func newNode(k int) *Node {
	return &Node{key: k}
}

func copyNode(n *Node) *Node {
	return &Node{
		key:   n.key,
		left:  n.left,
		right: n.right,
	}
}

// 向以node为根的二分搜索树中添加元素e
// 返回插入新节点后以node为根的二分搜索树
func (n *Node) add(k int) *Node {
	if n == nil {
		return newNode(k)
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
func (n *Node) contains(k int) bool {
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
func (n *Node) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.key)
	n.left.preOrder()
	n.right.preOrder()
}

// 以node为根节点的中序遍历
func (n *Node) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.key)
	n.right.inOrder()
}

func (n *Node) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.key)
}

// 获取以n为根节点的二分搜索树中的最小值
func (n *Node) minimum() *Node {
	if n.left == nil {
		return n
	}
	return n.left.minimum()
}

// 获取以n为根节点的二分搜索树中的最大值
func (n *Node) maximum() int {
	if n.right == nil {
		return n.key
	}
	return n.right.maximum()
}

// 删除以n为根节点的二分搜索树中的最小值
func (n *Node) removeMin() *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		return rightNode
	}
	n.left = n.left.removeMin()
	return n
}

// 删除以n为根节点的二分搜索树中的最小值
func (n *Node) removeMax() *Node {
	if n.right == nil {
		nodeLeft := n.left
		n.left = nil
		return nodeLeft
	}
	n.right = n.right.removeMax()
	return n
}

func (n *Node) remove(k int) *Node {
	if n == nil {
		return nil
	}
	// 先搜索k对应的节点
	if k < n.key {
		n.left = n.left.remove(k)
		return n
	} else if k > n.key {
		n.right = n.right.remove(k)
		return n
	} else { // k == n.key(找到要删除的节点了)
		// n的左孩子=nil，直接返回右孩子作为n
		if n.left == nil {
			return n.right
		}
		// 右孩子=nil
		if n.right == nil {
			return n.left
		}
		// 左右孩子都!=nil
		successor := copyNode(n.right.minimum())
		successor.right = n.right.removeMin()
		successor.left = n.left
		return successor
	}
}

func (n *Node) String(depth int) string {
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

func (n *Node) generateBSTString(depth int, res *strings.Builder) {
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