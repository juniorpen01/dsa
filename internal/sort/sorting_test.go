package sort_test

import (
	"slices"
	"testing"

	"github.com/juniorpen01/dsa/internal/sort"
)

func TestSorts(t *testing.T) {
	cases := []struct {
		name string
		sort func([]int) []int
	}{
		{"Bubble Sort", sort.BubbleSort},
		{"Merge Sort", sort.MergeSort},
		{"Insertion Sort", sort.InsertionSort},
		{"Selection Sort", sort.SelectionSort},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := []int{8, 3, 4, 7, 7, 2, 2, 1, 0, 0, 9, 6, 5}

			expected := []int{0, 0, 1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9}

			if actual := c.sort(input); !slices.Equal(actual, expected) {
				t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
			}
		})
	}

}
