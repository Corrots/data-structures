package v2

import (
	"fmt"
	"log"

	"github.com/corrots/data-structures/stack/linkedliststack"
)

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
	b.size++
}

func (b *BST) Contains(k int) bool {
	return b.root.contains(k)
}

// 二分搜索树的前序遍历
func (b *BST) PreOrder() {
	b.root.preOrder()
}

func (b *BST) PreOrderNR() {
	stack := linkedliststack.NewStack()
	stack.Push(b.root)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*Node)
		fmt.Printf("%d ", cur.key)
		if cur.right != nil {
			stack.Push(cur.right)
		}
		if cur.left != nil {
			stack.Push(cur.left)
		}
	}
}

// 二分搜索树的中序遍历
func (b *BST) InOrder() {
	b.root.inOrder()
}

// 二分搜索树的后序遍历
func (b *BST) PostOrder() {
	b.root.postOrder()
}

// 二分搜索树的层序遍历
func (b *BST) LevelOrder() {
	var q []*Node
	q = append(q, b.root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Printf("%d ", node.key)
		if node.left != nil {
			q = append(q, node.left)
		}
		if node.right != nil {
			q = append(q, node.right)
		}
	}
}

// 获取二分搜索树中的最小值
func (b *BST) Minimum() *Node {
	if b.IsEmpty() {
		log.Fatal("BST is empty")
	}
	return b.root.minimum()
}

// 获取二分搜索树中的最大值
func (b *BST) Maximum() int {
	if b.IsEmpty() {
		log.Fatal("BST is empty")
	}
	return b.root.maximum()
}

// 删除二分搜索树中的最小值
func (b *BST) RemoveMin() *Node {
	e := b.Minimum()
	b.root = b.root.removeMin()
	b.size--
	return e
}

// 删除二分搜索树中的最大值
func (b *BST) RemoveMax() int {
	e := b.Maximum()
	b.root = b.root.removeMax()
	b.size--
	return e
}

func (b *BST) Remove(k int) {
	if b.IsEmpty() {
		log.Fatal("BST is empty")
	}
	b.root = b.root.remove(k)
	return
}

func (b *BST) String() string {
	return b.root.String(0)
}
