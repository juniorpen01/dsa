package sort_test

import (
	"github.com/juniorpen01/dsa/internal/sort"
	"slices"
	"testing"
)

func TestSorts(t *testing.T) {
	cases := []struct {
		name string
		sort func([]int)
	}{
		{"Bubble Sort", sort.BubbleSort},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := []int{8, 3, 4, 7, 2, 1, 0, 9, 6, 5}

			expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

			actual := input

			if c.sort(actual); !slices.Equal(actual, expected) {
				t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
			}
		})
	}

}
