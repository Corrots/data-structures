package heap

import "log"

type MaxHeap struct {
	data []interface{}
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (h *MaxHeap) parent(index int) int {
	if index == 0 {
		log.Fatalf("invalid index,root doesn't have child node")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (h *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (h *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

// 向堆中添加元素
func (h *MaxHeap) Add(e interface{}) {
	h.data = append(h.data, e)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) siftUp(k int) {
	//  k不是根节点,且k的父亲节点<元素k
	for k > 0 && h.data[h.parent(k)].(int) < h.data[k].(int) {
		h.data[h.parent(k)], h.data[k] = h.data[k], h.data[h.parent(k)]
		k = h.parent(k)
	}
}

func (h *MaxHeap) FindMax() interface{} {
	n := len(h.data)
	if n == 0 {
		log.Fatal("heap is empty")
	}
	return h.data[0]
}

// 删除堆顶的元素
func (h *MaxHeap) ExtractMax() interface{} {
	length := len(h.data)
	if length == 0 {
		return nil
	}
	res := h.FindMax()
	h.data[0] = h.data[length-1]
	h.data = h.data[:length-1]
	h.siftDown(0)
	return res
}

func (h *MaxHeap) siftDown(k int) {
	n := len(h.data)
	// 循环中止条件:k的左孩子已越界
	for h.leftChild(k) < n {
		j := h.leftChild(k)
		if j+1 < n && h.data[j].(int) < h.data[j+1].(int) {
			j++
		}
		if h.data[k].(int) > h.data[j].(int) {
			break
		}
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}

// 取出堆顶的元素，并替换为元素e
func (h *MaxHeap) Replace(e interface{}) interface{} {
	res := h.FindMax()
	h.data[0] = e
	h.siftDown(0)
	return res
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}
