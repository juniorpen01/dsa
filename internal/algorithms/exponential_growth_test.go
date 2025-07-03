package algorithms_test

import (
	"slices"
	"testing"

	"github.com/juniorpen01/dsa/internal/algorithms"
)

func TestExponentialGrowth(t *testing.T) {
	cases := []struct {
		input struct {
			n, factor, days int
		}
		expected []int
	}{
		{struct {
			n, factor, days int
		}{10, 2, 4}, []int{10, 20, 40, 80, 160}},
	}

	for _, c := range cases {
		if actual, expected := algorithms.ExponentialGrowth(c.input.n, c.input.factor, c.input.days), c.expected; !slices.Equal(actual, expected) {
			t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
		}
	}
}
