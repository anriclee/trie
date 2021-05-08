package trie

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	trie := Constructor()
	trie.Insert("abc")
	trie.Insert("bad")
	trie.Insert("bae")
	trie.Insert("baf")

	//trie.Print()

	fmt.Println(trie.Search("abc") == true)
	fmt.Println(trie.Search("baf") == true)
}
