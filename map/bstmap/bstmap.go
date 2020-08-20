package bstmap

// 基于二分搜索树实现的Map
type BSTMap struct {
	bst *BST
}

func (m *BSTMap) Add(k int, v interface{}) {
	m.bst.Add(k, v)
}

func (m *BSTMap) Remove(k int) interface{} {
	return m.bst.Remove(k)
}

func (m *BSTMap) Contains(k int) bool {
	return m.bst.Contains(k)
}

func (m *BSTMap) Get(k int) interface{} {
	return m.bst.Get(k)
}

func (m *BSTMap) Set(k int, newVal interface{}) {
	m.bst.Set(k, newVal)
}

func (m *BSTMap) Len() int {
	return m.bst.Len()
}

func (m *BSTMap) IsEmpty() bool {
	return m.bst.IsEmpty()
}
