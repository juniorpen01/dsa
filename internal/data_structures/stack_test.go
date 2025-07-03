package datastructures_test

import (
	"slices"
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestStackPush(t *testing.T) {
	var stack datastructures.Stack
	stack.Push(1)
	stack.Push(2)
	if actual, expected := stack.Items, []int{1, 2}; !slices.Equal(actual, expected) {
		t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
	}
}

func TestStackSize(t *testing.T) {
	var stack datastructures.Stack
	stack.Push(1)
	stack.Push(2)
	if actual, expected := stack.Size(), 2; actual != expected {
		t.Errorf("Expected size does not match actual size: expected %v, got %v", expected, actual)
	}
}

func TestStackPeek(t *testing.T) {
	var stack datastructures.Stack
	stack.Push(1)
	stack.Push(2)
	expected := 2
	actual, ok := stack.Peek()
	if !ok {
		t.Errorf("Expected top item to exist: expected %t, got %t", true, ok)
	}
	if actual != expected {
		t.Errorf("Expected top item does not match actual top item: expected %d, got %d", expected, actual)
	}
}

func TestStackPop(t *testing.T) {

	var stack datastructures.Stack
	stack.Push(1)
	stack.Push(2)
	expected := 2
	actual, ok := stack.Pop()
	if !ok {
		t.Errorf("Expected top item to exist: expected %t, got %t", true, ok)
	}
	if actual != expected {
		t.Errorf("Expected top item does not match actual top item: expected %d, got %d", expected, actual)
	}
	if actual, expected := stack.Size(), 1; actual != expected {

		t.Errorf("Expected size does not match actual size: expected %d, got %d", expected, actual)
	}
}
