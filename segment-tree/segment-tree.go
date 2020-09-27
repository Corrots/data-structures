package segment_tree

import "log"

type SegmentTree struct {
	tree []int
	data []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	return &SegmentTree{
		tree: make([]int, n*4),
		data: arr,
	}
}

func buildSegmentTree(treeIndex, l, r int) {

}

func (st *SegmentTree) Get(i int) int {
	if i < 0 || i > len(st.data) {
		log.Fatalf("Invalid index: %d\n", i)
	}
	return st.data[i]
}

func (st *SegmentTree) Len() int {
	return len(st.data)
}

func (st *SegmentTree) leftChild(i int) int {
	return 2*i + 1
}

func (st *SegmentTree) rightChild(i int) int {
	return 2*i + 2
}
