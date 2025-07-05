package datastructures_test

import (
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestLinkedListAdd(t *testing.T) {
	cases := []struct {
		input  int
		action string
	}{
		{1, "tail"}, {2, "tail"}, {3, "tail"}, {4, "tail"}, {5, "tail"},
		{1, "head"}, {2, "head"}, {3, "head"}, {4, "head"}, {5, "head"},
	}

	var linkedList datastructures.LinkedList

	for _, c := range cases {
		switch c.action {
		case "head":
			linkedList.AddToHead(datastructures.NewNode(c.input))

			if actual, expected := linkedList.Head().Val(), c.input; actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
		case "tail":
			linkedList.AddToTail(datastructures.NewNode(c.input))

			var last datastructures.Node
			for cur := linkedList.Head(); cur != nil; cur = cur.Next() {
				last = *cur
			}

			if actual, expected := last.Val(), c.input; actual != expected {
				t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
			}
		}
	}
}

func TestLinkedListRemoveFromHead(t *testing.T) {
	cases := []struct {
		input int
	}{
		// {1, "tail"}, {2, "tail"}, {3, "tail"}, {4, "tail"}, {5, "tail"},
		{1}, {2}, {3}, {4}, {5},
	}

	var linkedList datastructures.LinkedList

	// Setup
	for _, c := range cases {
		linkedList.AddToTail(datastructures.NewNode(c.input))
	}

	// Test
	cur := linkedList.RemoveFromHead()

	for _, c := range cases {
		if actual, expected := cur.Val(), c.input; actual != expected {
			t.Errorf("Expected value does not match actual value: expected %d, got %d", expected, actual)
		}
		cur = linkedList.RemoveFromHead()
	}
}
