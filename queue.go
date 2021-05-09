package trie

type Queue struct {
	Nodes []*Trie
}

func NewQueue() *Queue {
	return &Queue{
		Nodes: []*Trie{},
	}
}

func (q *Queue) Enqueue(newNode *Trie) bool {
	if newNode == nil {
		return false
	}
	q.Nodes = append(q.Nodes, newNode)
	return true
}

func (q *Queue) Dequeue() *Trie {
	if len(q.Nodes) == 0 {
		return nil
	}
	ret := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return ret
}

func (q *Queue) Poll() *Trie {
	if len(q.Nodes) == 0 {
		return nil
	}
	return q.Nodes[0]
}
