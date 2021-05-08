package trie

import "fmt"

type Trie struct {
	Children [26]*Trie
	Val      byte
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	root := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if root.Children[idx] == nil {
			newNode := Constructor()
			newNode.Val = word[i]
			root.Children[idx] = &newNode
		}
		root = root.Children[idx]
	}
	root.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	ptr := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if ptr.Children[idx] == nil {
			return false
		}
		ptr = ptr.Children[idx]
	}
	return ptr.IsEnd
}

func (t *Trie) Print() {
	// 后序遍历验证是否正确
	for _, child := range t.Children {
		if child.Val != 0 {
			child.Print()
			fmt.Println(string(child.Val))
		}
	}
}
