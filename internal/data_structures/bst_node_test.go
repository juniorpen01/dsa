package datastructures_test

import (
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestBSTNodeInsert(t *testing.T) {
	cases := struct {
		input    []int
		expected struct {
			left, right []int
		}
	}{
		[]int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5},
		struct {
			left  []int
			right []int
		}{
			[]int{0, -1, -2, -3, -4, -5},
			[]int{0, 1, 2, 3, 4, 5},
		},
	}

	root := datastructures.NewBSTNode(0)

	for _, c := range cases.input {
		root.Insert(c)
	}

	// Test Left
	cur := &root
	for _, expected := range cases.expected.left {
		if actual := cur.Val(); actual != expected {
			t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
		}
		cur = cur.Left()
	}

	// Test Right
	cur = &root
	for _, expected := range cases.expected.right {
		if actual := cur.Val(); actual != expected {
			t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
		}
		cur = cur.Right()
	}
}
