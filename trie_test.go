package trie

import (
	"fmt"
	"testing"
)

func TestTransitionArray(t *testing.T) {
	trie := New(1, '0')
	//ace add bad bed bee cab dad
	trie.Insert("ace")
	trie.Insert("add")
	trie.Insert("bad")
	trie.Insert("bed")
	trie.Insert("bee")
	trie.Insert("cab")
	trie.Insert("dad")

	matrix := trie.TransitionMatrix()
	TrimAndPrintMatrix(matrix)

	//fmt.Println(trie.Search("ace") == true)
	//fmt.Println(trie.Search("bee") == true)
	//fmt.Println(trie.Search("bef") == false)

	fmt.Println(SearchByMatrix(matrix, "ace") == true)
	fmt.Println(SearchByMatrix(matrix, "bee") == true)
	fmt.Println(SearchByMatrix(matrix, "bef") == false)
}
