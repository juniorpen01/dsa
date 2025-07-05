package datastructures

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) Node {
	return Node{val, nil}
}

func (n *Node) Val() int {
	return n.val
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(node Node) {
	n.next = &node
}
