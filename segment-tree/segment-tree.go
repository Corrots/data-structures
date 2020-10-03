package segment_tree

import (
	"fmt"
	"log"
	"strings"
)

type SegmentTree struct {
	tree   []int
	data   []int
	merger func(v1, v2 int) int
}

func NewSegmentTree(arr []int, mergeFunc func(int, int) int) *SegmentTree {
	n := len(arr)
	st := &SegmentTree{
		tree:   make([]int, n*4),
		data:   arr,
		merger: mergeFunc,
	}
	st.buildSegmentTree(0, 0, n-1)
	return st
}

// 在treeIndex位置创建表示区间[l...r]的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	mid := (r-l)/2 + l
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
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

func (st *SegmentTree) String() string {
	var res strings.Builder
	res.WriteString("[")
	for i := 0; i < len(st.tree); i++ {
		res.WriteString(fmt.Sprintf("%d ", st.tree[i]))
	}
	res.WriteString("]")
	return res.String()
}
