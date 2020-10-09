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
	return &Node{isWord: isWord}
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Len() int {
	return t.size
}

// 向Trie中添加新单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.next[c] == nil {
			cur.next[c] = newNode(false)
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}
