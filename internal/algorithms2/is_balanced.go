package algorithms2

import datastructures "github.com/juniorpen01/dsa/internal/data_structures"

func IsBalanced(str string) bool {
	// // HACK: Technically this whole thing, apparently I was supposed to use a stack
	// counter := 0
	// for _, c := range str {
	// 	switch c {
	// 	case '(':
	// 		counter++
	// 	case ')':
	// 		counter--
	// 	}
	// 	if counter < 0 {
	// 		return false
	// 	}
	// }
	// return counter == 0

	var stack datastructures.Stack
	for _, c := range str {
		switch c {
		case '(':
			stack.Push(int(c)) // HACK: Kinda hacky (and does come from hacky solution) but casting rune to int works
		case ')':
			if _, ok := stack.Pop(); !ok {
				return false
			}
		}
	}
	return stack.Size() == 0
}
