package v1

type Node struct {
	isEnd bool
	next  map[uint8]*Node
}

type Trie struct {
	root *Node
	size int
}

func newNode(isEnd bool) *Node {
	return &Node{
		isEnd: isEnd,
		next:  make(map[uint8]*Node),
	}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: newNode(false)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode(false)
		}
		cur = cur.next[c]
	}
	if !cur.isEnd {
		cur.isEnd = true
		this.size++
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
