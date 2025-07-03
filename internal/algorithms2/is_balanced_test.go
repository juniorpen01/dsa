package algorithms2_test

import (
	"testing"

	"github.com/juniorpen01/dsa/internal/algorithms2"
)

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"(", false},
		{"()", true},
		{"(())", true},
		{"()()", true},
		{"(()))", false},
		{"((())())", true},
		{"(()(()", false},
		{")(", false},
		{")()(()", false},
		{"())(()", false},

		{"((())(()))", true},
	}

	for _, c := range cases {
		if actual, expected := algorithms2.IsBalanced(c.input), c.expected; actual != expected {
			t.Errorf("Expected does not match actual: input %s, expected %v, got %v", c.input, expected, actual)
		}
	}
}
