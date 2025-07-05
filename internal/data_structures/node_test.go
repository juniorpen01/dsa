package datastructures_test

import (
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestNode(t *testing.T) {
	cases := []struct {
		input    int
		expected []int
	}{
		{1, []int{0, 1}},
		{2, []int{0, 1, 2}},
		{3, []int{0, 1, 2, 3}},
		{4, []int{0, 1, 2, 3, 4}},
		{5, []int{0, 1, 2, 3, 4, 5}},
	}

	linkedList := datastructures.NewNode(0)
	head := &linkedList

	for _, c := range cases {
		head.SetNext(datastructures.NewNode(c.input))
		head = head.Next()

		// Test
		current := &linkedList
		for i := range len(c.expected) {
			if actual, expected := current.Val(), c.expected[i]; actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
			current = current.Next()
		}

	}
}
