package heap

import "log"

// 最小堆
type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// 返回数组表示的完全二叉树中k对应的父节点的索引
func (h *MinHeap) parent(k int) int {
	if k == 0 {
		log.Fatalln("invalid index, root doesn't have child node")
	}
	return (k - 1) / 2
}

// 返回数组表示的完全二叉树中k对应的左孩子的索引
func (h *MinHeap) leftChild(k int) int {
	return 2*k + 1
}

// 返回数组表示的完全二叉树中k对应的左孩子的索引
func (h *MinHeap) rightChild(k int) int {
	return 2*k + 2
}

// 向最小堆中添加元素e
func (h *MinHeap) Add(e int) {
	h.data = append(h.data, e)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) siftUp(k int) {
	for k > 0 && h.data[h.parent(k)] > h.data[k] {
		h.data[h.parent(k)], h.data[k] = h.data[k], h.data[h.parent(k)]
		k = h.parent(k)
	}
}

// 找到并返回堆中的最小值(堆顶元素)
func (h *MinHeap) FindMin() int {
	if len(h.data) == 0 {
		log.Fatal("min heap is empty")
	}
	return h.data[0]
}

// 推出堆中的最小值(堆顶元素)
func (h *MinHeap) ExtractMin() int {
	res := h.FindMin()
	n := len(h.data)
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.siftDown(0)
	return res
}

func (h *MinHeap) siftDown(k int) {
	n := len(h.data)
	// 当k的左孩子越界时，循环终止
	for h.leftChild(k) < n {
		j := h.leftChild(k)
		if j+1 < n && h.data[j+1] < h.data[j] {
			j++
		}
		// j记录的是左右孩子中的最小值索引
		// 若节点k的值已经<=左右孩子节点的值，则循环终止
		if h.data[k] <= h.data[j] {
			break
		}
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}

// 取出堆顶的元素，并替换为e
func (h *MinHeap) Replace(e int) int {
	res := h.FindMin()
	h.data[0] = e
	h.siftDown(0)
	return res
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}
