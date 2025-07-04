package datastructures_test

import (
	"slices"
	"testing"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestQueuePush(t *testing.T) {
	var Queue datastructures.Queue
	Queue.Push(1)
	Queue.Push(2)
	if actual, expected := Queue.Items, []int{2, 1}; !slices.Equal(actual, expected) {
		t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
	}
}

func TestQueueSize(t *testing.T) {
	var Queue datastructures.Queue
	Queue.Push(1)
	Queue.Push(2)
	if actual, expected := Queue.Size(), 2; actual != expected {
		t.Errorf("Expected size does not match actual size: expected %v, got %v", expected, actual)
	}
}

func TestQueuePeek(t *testing.T) {
	var Queue datastructures.Queue
	Queue.Push(1)
	Queue.Push(2)
	expected := 1
	actual, ok := Queue.Peek()
	if !ok {
		t.Errorf("Expected top item to exist: expected %t, got %t", true, ok)
	}
	if actual != expected {
		t.Errorf("Expected top item does not match actual top item: expected %d, got %d", expected, actual)
	}
}

func TestQueuePop(t *testing.T) {

	var Queue datastructures.Queue
	Queue.Push(1)
	Queue.Push(2)
	expected := 1
	actual, ok := Queue.Pop()
	if !ok {
		t.Errorf("Expected top item to exist: expected %t, got %t", true, ok)
	}
	if actual != expected {
		t.Errorf("Expected top item does not match actual top item: expected %d, got %d", expected, actual)
	}
	if actual, expected := Queue.Size(), 1; actual != expected {

		t.Errorf("Expected size does not match actual size: expected %d, got %d", expected, actual)
	}
}
