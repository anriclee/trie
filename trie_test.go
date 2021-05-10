package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	require.True(t, SearchByMatrix(matrix, "ace"))
	require.True(t, SearchByMatrix(matrix, "add"))
	require.True(t, SearchByMatrix(matrix, "bad"))
	require.True(t, SearchByMatrix(matrix, "bed"))
	require.True(t, SearchByMatrix(matrix, "bee"))
	require.True(t, SearchByMatrix(matrix, "cab"))
	require.True(t, SearchByMatrix(matrix, "dad"))

	require.False(t, SearchByMatrix(matrix, "ac"))
	require.False(t, SearchByMatrix(matrix, "ad"))
	require.False(t, SearchByMatrix(matrix, "ba"))
	require.False(t, SearchByMatrix(matrix, "bef"))
	require.False(t, SearchByMatrix(matrix, "beee"))
}
