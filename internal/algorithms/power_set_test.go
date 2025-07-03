package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/juniorpen01/dsa/internal/algorithms"
)

func TestPowerSet(t *testing.T) {
	cases := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, c := range cases {
		if actual, expected := algorithms.PowerSet(c.input), c.expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
		}
	}
}
