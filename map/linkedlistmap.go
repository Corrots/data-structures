package _map

// 基于linkedlist实现的Map
type LinkedListMap struct {
	list *LinkedList
}

func (m *LinkedListMap) Add(k int, v interface{}) {
	m.list.Add(k, v)
}

func (m *LinkedListMap) Remove(i int) interface{} {
	return m.list.Remove(i)
}

func (m *LinkedListMap) Contains(i int) bool {
	return m.list.Contains(i)
}

func (m *LinkedListMap) Get(i int) interface{} {
	return m.list.Get(i)
}

func (m *LinkedListMap) Set(k int, v interface{}) {
	m.list.Set(k, v)
}

func (m *LinkedListMap) Len() int {
	return m.list.Len()
}

func (m *LinkedListMap) IsEmpty() bool {
	return m.list.IsEmpty()
}
