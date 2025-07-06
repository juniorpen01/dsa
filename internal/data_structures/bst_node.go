package datastructures

type BSTNode struct {
	val         int
	left, right *BSTNode
}

func NewBSTNode(val int) BSTNode {
	return BSTNode{val: val}
}

func (b *BSTNode) Insert(val int) {
	switch {
	case b.val > val:
		if b.left != nil {
			b.left.Insert(val)
		} else {
			left := NewBSTNode(val)
			b.left = &left
		}
	case b.val < val:
		if b.right != nil {
			b.right.Insert(val)
		} else {
			right := NewBSTNode(val)
			b.right = &right
		}
	}
}

func (b *BSTNode) Min() int {
	min := 0
	for cur := b; cur != nil; cur = cur.left {
		min = cur.val
	}
	return min
}

func (b *BSTNode) Max() int {
	max := 0
	for cur := b; cur != nil; cur = cur.right {
		max = cur.val
	}
	return max
}

func (b *BSTNode) Val() int {
	return b.val
}

func (b *BSTNode) Left() *BSTNode {
	return b.left
}

func (b *BSTNode) Right() *BSTNode {
	return b.right
}
