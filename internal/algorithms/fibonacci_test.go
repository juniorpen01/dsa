package algorithms_test

import (
	"testing"

	"github.com/juniorpen01/dsa/internal/algorithms"
)

func TestFIbonacci(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{45, 1134903170},
	}

	for _, c := range cases {
		if actual, expected := algorithms.Fibonacci(c.input), c.expected; actual != expected {
			t.Errorf("Expected output does not match actual output: expected %d, got %d", expected, actual)
		}
	}
}
