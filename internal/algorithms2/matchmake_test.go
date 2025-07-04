package algorithms2_test

import (
	"slices"
	"testing"

	"github.com/juniorpen01/dsa/internal/algorithms2"
	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

func TestMatchmake(t *testing.T) {
	type Expected struct {
		data    []int
		message string
	}
	cases := []struct {
		input    algorithms2.User
		expected Expected
	}{
		{algorithms2.User{0, "join"}, Expected{[]int{0}, "No match found"}},
		{algorithms2.User{1, "join"}, Expected{[]int{1, 0}, "No match found"}},
		{algorithms2.User{2, "join"}, Expected{[]int{2, 1, 0}, "No match found"}},
		{algorithms2.User{3, "join"}, Expected{[]int{3, 2}, "0 matched 1!"}},
		{algorithms2.User{4, "join"}, Expected{[]int{4, 3, 2}, "No match found"}},
		{algorithms2.User{5, "join"}, Expected{[]int{5, 4}, "2 matched 3!"}},
		{algorithms2.User{5, "leave"}, Expected{[]int{4}, "No match found"}},
		{algorithms2.User{4, "leave"}, Expected{[]int{}, "No match found"}},
		{algorithms2.User{6, "join"}, Expected{[]int{6}, "No match found"}},
		{algorithms2.User{6, "leave"}, Expected{[]int{}, "No match found"}},
		{algorithms2.User{7, "join"}, Expected{[]int{7}, "No match found"}},
		{algorithms2.User{8, "join"}, Expected{[]int{8, 7}, "No match found"}},
		{algorithms2.User{9, "join"}, Expected{[]int{9, 8, 7}, "No match found"}},
		{algorithms2.User{10, "join"}, Expected{[]int{10, 9}, "7 matched 8!"}},
	}

	var queue datastructures.Queue

	for _, c := range cases {
		message := algorithms2.Matchmake(&queue, c.input)
		if actual, expected := message, c.expected.message; message != expected {
			t.Errorf("Expected message does not match actual message: expected %v, got %v", expected, actual)
		}
		if actual, expected := queue.Items, c.expected.data; !slices.Equal(actual, expected) {
			t.Errorf("Expected data does not match actual data: expected %v, got %v", expected, actual)
		}
	}
}
