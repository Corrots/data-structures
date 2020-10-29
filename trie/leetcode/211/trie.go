package v1

type Node struct {
	isEnd bool
	next  map[uint8]*Node
}

type WordDictionary struct {
	root *Node
}

func newNode(isEnd bool) *Node {
	return &Node{
		isEnd: isEnd,
		next:  make(map[uint8]*Node),
	}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: newNode(false)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode(false)
		}
		cur = cur.next[c]
	}
	cur.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match(this.root, word, 0)
}

func (this *WordDictionary) match(node *Node, word string, index int) bool {
	if index == len(word) {
		return node.isEnd
	}
	c := word[index]
	if c != '.' {
		if _, ok := node.next[c]; !ok {
			return false
		}
		return this.match(node.next[c], word, index+1)
	}
	// 若c=='.' 遍历下一个字符所有可能
	for k := range node.next {
		if this.match(node.next[k], word, index+1) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
