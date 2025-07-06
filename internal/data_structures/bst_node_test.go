package datastructures_test

import (
	"slices"
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

func TestBSTNodeMax(t *testing.T) {
	cases := struct {
		input    []int
		expected int
	}{
		[]int{-2, 1, -3, 4, 2, -5, 5, -4, -1, 3}, 5,
	}

	root := datastructures.NewBSTNode(0)

	for _, c := range cases.input {
		root.Insert(c)
	}

	if actual, expected := root.Max(), cases.expected; actual != expected {
		t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
	}
}

func TestBSTNodeMin(t *testing.T) {
	cases := struct {
		input    []int
		expected int
	}{
		[]int{-2, 1, -3, 4, 2, -5, 5, -4, -1, 3}, -5,
	}

	root := datastructures.NewBSTNode(0)

	for _, c := range cases.input {
		root.Insert(c)
	}

	if actual, expected := root.Min(), cases.expected; actual != expected {
		t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
	}
}

func TestBSTNodeDelete(t *testing.T) {
	cases := []struct {
		input    int
		expected struct {
			left, right []int
		}
	}{
		{
			2, struct {
				left  []int
				right []int
			}{[]int{0, -1, -2, -3, -4, -5}, []int{0, 1, 2}}}, {-3, struct {
			left  []int
			right []int
		}{[]int{0, -1, -2}, []int{0, 1, 2}}}, // HACK: Shut UP!!!
	}

	root := datastructures.NewBSTNode(0)

	for _, val := range []int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5} {
		root.Insert(val)
	}

	for _, c := range cases {
		root.Delete(c.input)

		// Test Left
		cur := &root
		for _, expected := range c.expected.left {
			if actual := cur.Val(); actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
			cur = cur.Left()
		}

		// Test Right
		cur = &root
		for _, expected := range c.expected.right {
			if actual := cur.Val(); actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
			cur = cur.Right()
		}
	}
}

func TestBSTNodePreorder(t *testing.T) {
	cases := []struct {
		input, expected []int
	}{
		{[]int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5}, []int{0, -1, -2, -3, -4, -5, 1, 2, 3, 4, 5}},
		{[]int{-2, 1, -3, 4, 2, -5, 5, -4, -1, 3}, []int{0, -2, -3, -5, -4, -1, 1, 4, 2, 3, 5}},
	}

	for _, c := range cases {
		root := datastructures.NewBSTNode(0)
		for _, val := range c.input {
			root.Insert(val)
		}

		var arr []int
		if actual, expected := root.Preorder(&arr), c.expected; !slices.Equal(actual, expected) {
			t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
		}
	}
}

func TestBSTNodePostorder(t *testing.T) {
	cases := []struct {
		input, expected []int
	}{
		{[]int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{[]int{-2, 1, -3, 4, 2, -5, 5, -4, -1, 3}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		root := datastructures.NewBSTNode(0)
		for _, val := range c.input {
			root.Insert(val)
		}

		var arr []int
		if actual, expected := root.Inorder(&arr), c.expected; !slices.Equal(actual, expected) {
			t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
		}
	}
}

func TestBSTNodeExists(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{{-3, true}, {2, true}, {6, false}, {727, false}}

	root := datastructures.NewBSTNode(0)

	for _, val := range []int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5} {
		root.Insert(val)
	}

	for _, c := range cases {
		if actual, expected := root.Exists(c.input), c.expected; actual != expected {
			t.Errorf("Expected does not match actual: expected %t, got %t", expected, actual)
		}
	}
}

func TestBSTNodeHeight(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{-1, 1, -2, 2, 3, -3, 4, -4, -5, 5}, 6},
		{[]int{-2, 1, -3, 4, 2, -5, 5, -4, -1, 3}, 5},
	}

	for _, c := range cases {
		root := datastructures.NewBSTNode(0)

		for _, val := range c.input {
			root.Insert(val)
		}

		if actual, expected := root.Height(), c.expected; actual != expected {
			t.Errorf("Expected height does not match actual height: expected %d, got %d", expected, actual)
		}
	}

}
