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

func (b *BSTNode) Delete(val int) *BSTNode {
	switch {
	case b.val > val:
		b.left.Delete(val)
	case b.val < val:
		b.right.Delete(val)
	default:
		switch {
		case b.left == nil:
			return b.right
		case b.right == nil:
			return b.left
		default:
			min := b.right.Min()
			b.val = min
			b.right.Delete(min)
		}
	}

	return b
}

func (b *BSTNode) Preorder(visited *[]int) []int {
	if b == nil {
		return *visited
	}
	*visited = append(*visited, b.val)
	b.left.Preorder(visited)
	b.right.Preorder(visited)
	return *visited
}

func (b *BSTNode) Postorder(visited *[]int) []int {
	if b == nil {
		return *visited
	}
	b.left.Postorder(visited)
	b.right.Postorder(visited)
	*visited = append(*visited, b.val)
	return *visited
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
