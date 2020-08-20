package bstmap

import (
	"log"
)

type BST struct {
	root *TreeNode
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

func (b *BST) Get(k int) interface{} {
	node := b.root.getTreeNode(k)
	if node != nil {
		return node.val
	}
	return nil
}

func (b *BST) Set(k int, val interface{}) {
	node := b.root.getTreeNode(k)
	if node == nil {
		log.Fatalf("key: %d doesn't exist\n", k)
	}
	node.val = val
}

// 向二分搜索树中插入元素
func (b *BST) Add(k int, v interface{}) {
	b.root = b.root.add(k, v)
	b.size++
}

func (b *BST) Contains(k int) bool {
	return b.root.getTreeNode(k) != nil
}

// 获取二分搜索树中的最小值
func (b *BST) Minimum() *TreeNode {
	if b.IsEmpty() {
		log.Fatal("BST is empty")
	}
	return b.root.minimum()
}

// 删除二分搜索树中的最小值
func (b *BST) RemoveMin() *TreeNode {
	e := b.Minimum()
	b.root = b.root.removeMin()
	b.size--
	return e
}

func (b *BST) Remove(k int) interface{} {
	node := b.root.getTreeNode(k)
	if node != nil {
		b.root = b.root.remove(k)
		return node.val
	}
	return nil
}
