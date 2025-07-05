package datastructures

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddToHead(node Node) {
	node.SetNext(*l.Head())
	l.SetHead(node)
}

func (l *LinkedList) RemoveFromHead() *Node {
	if l.Head() == nil {
		return nil
	}

	poppedNode := *l.Head()

	if l.Head().Next() == nil {
		l.head = nil
	} else {
		l.SetHead(*l.Head().Next())
	}

	return &poppedNode
}

func (l *LinkedList) AddToTail(node Node) {
	if l.Head() == nil {
		l.SetHead(node)
		return
	}
	var last *Node
	for cur := l.Head(); cur != nil; cur = cur.Next() {
		last = cur
	}
	last.SetNext(node)
}

func (l LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) SetHead(node Node) {
	l.head = &node
}
