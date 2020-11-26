package fib

// Fib returns fibonacci number by given offset
func Fib(n int) int {
	a := 0
	b := 1
	if n > 0 {
		for n > 1 {
			n--
			t := a + b
			a = b
			b = t
		}
		return b
	}
	return 0
}
