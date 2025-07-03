package algorithms

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	grandparent, parent := 0, 1
	var current int
	for i := 0; i < n-1; i++ {
		current = grandparent + parent
		grandparent = parent
		parent = current
	}
	return current
}
