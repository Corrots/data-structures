package heap

import "log"

type MaxHeap struct {
	data []interface{}
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (h *MaxHeap) Parent(index int) int {
	if index == 0 {
		log.Fatalf("invalid index,root doesn't have child node")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (h *MaxHeap) LeftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (h *MaxHeap) RightChild(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) Add(e interface{}) {
	h.data = append(h.data, e)
	h.shiftUp(len(h.data) - 1)
}

// 向堆中添加元素
func (h *MaxHeap) shiftUp(k int) {
	// 当k不为根节点,且k的父亲节点<元素k
	for k > 0 && h.data[h.Parent(k)] < h.data[k] {
		h.data[h.Parent(k)], h.data[k] = h.data[k], h.data[h.Parent(k)]
		k = h.Parent(k)
	}
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}
