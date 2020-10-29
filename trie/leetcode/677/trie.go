package v1

type Node struct {
	value int
	next  map[uint8]*Node
}

type MapSum struct {
	root *Node
}

func newNode(val int) *Node {
	return &Node{
		value: val,
		next:  make(map[uint8]*Node),
	}
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{root: newNode(0)}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for i := 0; i < len(key); i++ {
		c := key[i]
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode(0)
		}
		cur = cur.next[c]
	}
	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := cur.next[c]; !ok {
			return 0
		}
		cur = cur.next[c]
	}
	// 遍历以prefix开头的所有字符串
	return sum(cur)
}

func sum(node *Node) int {
	res := node.value
	for k := range node.next {
		if n, ok := node.next[k]; ok {
			res += sum(n)
		}
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
