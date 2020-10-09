package v2

import (
	"fmt"
	"log"
	"strings"
)

type NumArray struct {
	Tree *SegmentTree
}

func Constructor(nums []int) NumArray {
	return NumArray{Tree: NewSegmentTree(nums, sumFunc)}
}

func (this *NumArray) Update(i int, val int) {
	if this == nil {
		log.Fatal("NumArray is nil")
	}
	if i < 0 || i >= this.Tree.Len() {
		log.Fatalf("invalid index %d\n", i)
	}
	this.Tree.Set(i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	if this == nil {
		log.Fatal("NumArray is nil")
	}
	return this.Tree.Query(i, j)
}

func sumFunc(i, j int) int {
	return i + j
}

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
	return st.tree[i]
}

// 在以rootIndex为根的线段树中更新index的值为val
func (st *SegmentTree) Set(index, val int) {
	if index < 0 || index > st.Len() {
		log.Fatal("invalid index")
	}
	st.set(0, 0, st.Len()-1, index, val)
}

// 在以treeIndex为根的线段树中，更新index的值为val
func (st *SegmentTree) set(treeIndex, l, r, index, val int) {
	if l == r {
		st.tree[treeIndex] = val
		return
	}
	mid := (r-l)/2 + l
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	if index >= mid+1 {
		st.set(rightTreeIndex, mid+1, r, index, val)
	} else { // index <= mid
		st.set(leftTreeIndex, l, mid, index, val)
	}

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// 返回区间[queryL, queryR]的值
func (st *SegmentTree) Query(queryL, queryR int) int {
	if queryL < 0 || queryL >= st.Len() || queryR < 0 || queryR >= st.Len() || queryL > queryR {
		log.Fatal("invalid index")
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以treeID为根的线段树中[l, r]范围内，搜索区间[queryL, queryR]的值
func (st *SegmentTree) query(treeIndex, l, r, queryL, queryR int) int {
	if queryL == l && queryR == r {
		return st.tree[treeIndex]
	}
	mid := (r-l)/2 + l
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	if queryL >= mid+1 {
		// 结果全部分布在右区间
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		// 结果全部分布在左区间
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	}
	// 结果分布在左右两个区间
	leftRes := st.query(leftTreeIndex, l, mid, queryL, mid)
	rightRes := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)
	return st.merger(leftRes, rightRes)
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

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
