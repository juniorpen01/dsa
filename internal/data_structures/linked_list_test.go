package datastructures_test

import (
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestNodePush(t *testing.T) {
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
	newNode := &linkedList

	for _, c := range cases {
		newNode.SetNext(datastructures.NewNode(c.input))
		newNode = newNode.Next()

		// Test
		currentNode := &linkedList
		for i := range len(c.expected) {
			if actual, expected := currentNode.Val(), c.expected[i]; actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
			if currentNode.Next() == nil {
				break
			} else {
				currentNode = currentNode.Next()
			}
		}

	}

	for true {
		t.Log(linkedList.Val())
		if linkedList.Next() == nil {
			break
		}
		linkedList = *linkedList.Next()
	}
}
