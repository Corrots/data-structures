package trie

type Node struct {
	isWord bool
	next   map[uint8]*Node
}

type Trie struct {
	root *Node
	size int
}

func newNode(isWord bool) *Node {
	return &Node{
		isWord: isWord,
		next:   make(map[uint8]*Node),
	}
}

func NewTrie() *Trie {
	return &Trie{root: newNode(false)}
}

func (t *Trie) Len() int {
	return t.size
}

// 向Trie中添加新单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode(false)
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

// 查询单词word是否存在Trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}
