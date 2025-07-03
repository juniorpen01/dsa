package other_test

import (
	"testing"

	"github.com/juniorpen01/dsa/internal/other"
)

func TestCountMarketers(t *testing.T) {
	cases := []struct {
		input    []string
		expected int
	}{
		{[]string{"developer", "marketer", "designer"}, 1},
		{[]string{"marketer", "marketer", "developer", "marketer"}, 3},
		{[]string{}, 0},
		{[]string{"developer", "designer", "product manager"}, 0},
		{[]string{"marketer"}, 1},
		{[]string{"marketer", "Marketer", "MARKETER"}, 3},
	}

	for _, c := range cases {
		if actual, expected := other.CountMarketers(c.input), c.expected; actual != expected {
			t.Errorf("Expected count does not match actual count: expected %d, got %d", expected, actual)
		}
	}
}
