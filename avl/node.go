package avl

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	key int
	val interface{}
	// 当前节点所处的高度，初始为1
	height int
	left   *TreeNode
	right  *TreeNode
}

func newTreeNode(k int, val interface{}) *TreeNode {
	// 初始化节点高度为1
	return &TreeNode{key: k, val: val, height: 1}
}

// 判断以n为根的二叉树是否为平衡二叉树
func (n *TreeNode) isBalanced() bool {
	if n == nil {
		return true
	}
	if abs(n.getBalanceFactor()) > 1 {
		return false
	}
	return n.left.isBalanced() && n.right.isBalanced()
}

// 获取节点n对应的高度
func (n *TreeNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// 计算节点n的平衡因子 = 左子树的高度-右子树的高度
func (n *TreeNode) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

// 向以node为根的二分搜索树中添加元素e
// 返回插入新节点后以node为根的二分搜索树
func (n *TreeNode) add(k int, val interface{}) *TreeNode {
	if n == nil {
		return newTreeNode(k, val)
	}
	// k == n.key 不做任何操作
	if k < n.key {
		n.left = n.left.add(k, val)
	} else if k > n.key {
		n.right = n.right.add(k, val)
	}
	// 更新height
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	// 计算平衡因子
	balanceFactor := n.getBalanceFactor()
	if abs(balanceFactor) > 1 {
		fmt.Println("unbalanced " + string(balanceFactor))
	}
	// 平衡维护
	if balanceFactor > 1 && n.left.getBalanceFactor() >= 0 {

	}

	return n
}

// 在以n为根节点的二分搜索树中查找k所在的节点
func (n *TreeNode) getNode(k int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.key == k {
		return n
	} else if n.key < k {
		return n.right.getNode(k)
	} else {
		return n.left.getNode(k)
	}
}

// 查找以node为根的二分搜索树中是否包含节点k
func (n *TreeNode) contains(k int) bool {
	return n.getNode(k) != nil
}

// 获取以n为根节点的二分搜索树中的最小值
func (n *TreeNode) minimum() *TreeNode {
	if n.left == nil {
		return n
	}
	return n.left.minimum()
}

// 获取以n为根节点的二分搜索树中的最大值
func (n *TreeNode) maximum() *TreeNode {
	if n.right == nil {
		return n
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

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
