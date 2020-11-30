package fib

// Fib returns fibonacci number by given offset
func Fib(n int) int {
	a, b := 0, 1
	if n > 0 {
		for n > 1 {
			n--
			a, b = b, a+b
		}
		return b
	}
	return 0
}
