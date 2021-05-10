package trie

import "fmt"

const size = 27
const terminateLabel = '#'

var (
	nodeVal       = 1
	terminateNode *Trie
	maxChar       = uint8('a')
)

type Trie struct {
	Children [size]*Trie
	Value    int
	Label    byte
	IsEnd    bool
}

func New(value int, label byte) Trie {
	return Trie{
		Value: value,
		Label: label,
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	root := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if word[i] > maxChar {
			maxChar = word[i]
		}
		if root.Children[idx] == nil {
			nodeVal++
			newNode := New(nodeVal, word[i])
			root.Children[idx] = &newNode
		}
		root = root.Children[idx]
	}
	root.IsEnd = true
	if terminateNode == nil {
		nodeVal++
		terminateNode = &Trie{
			Label: terminateLabel,
			Value: nodeVal,
		}
	}
	root.Children[size-1] = terminateNode
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

func (t *Trie) TransitionMatrix() [][]int {
	matrix := newMatrix()
	root := t
	q := NewQueue()
	q.Enqueue(root)
	for q.Poll() != nil {
		node := q.Dequeue()
		fmt.Println("label", string(node.Label), ":", node.Value)
		for _, child := range node.Children {
			if child == nil {
				continue
			}
			if child.Label == terminateLabel {
				matrix[node.Value][size-1] = child.Value
			} else {
				matrix[node.Value][child.Label-'a'] = child.Value
			}
			q.Enqueue(child)
		}
	}
	return matrix
}

func TrimAndPrintMatrix(matrix [][]int) [][]int {
	fmt.Println("transition matrix as follows:")
	trimFromCol := int(maxChar-'a') + 1
	trimToCol := 25
	fmt.Printf("%4s", "#st")
	for col := 0; col < len(matrix[0])-1; col++ {
		if col >= trimFromCol && col <= trimToCol {
			continue
		}
		fmt.Printf("%4s", string(byte('a'+col)))
	}
	fmt.Printf("%4s", "EOF")
	fmt.Println()
	for row := 0; row < len(matrix); row++ {
		rowPart1 := matrix[row][:trimFromCol]
		rowPart2 := matrix[row][trimToCol+1:]
		matrix[row] = append(rowPart1, rowPart2...)
	}
	for row := 0; row < len(matrix); row++ {
		fmt.Printf("%4d", row)
		for col := 0; col < len(matrix[0]); col++ {
			fmt.Printf("%4d", matrix[row][col])
		}
		fmt.Println()
	}
	return matrix
}

func SearchByMatrix(matrix [][]int, word string) bool {
	rowCnt := len(matrix)
	if rowCnt < 2 {
		return false
	}
	row := matrix[1]
	var i int
	for ; i < len(word); i++ {
		col := int(word[i] - 'a')
		if col >= len(matrix[1])-1 {
			return false
		}
		row = matrix[row[col]]
	}
	// 最后遍历到的一行可以直接到终结节点，才算匹配
	eof := row[len(row)-1]
	return eof != 0
}

func newMatrix() [][]int {
	matrix := make([][]int, 20)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, size)
	}
	return matrix
}

func (t *Trie) Print() {
	// 后序遍历验证是否正确
	for _, child := range t.Children {
		if child != nil {
			child.Print()
			fmt.Println("label", string(child.Label), ":", child.Value)
		}
	}
}
