package math

// Calculates and returns factorial of n.
func Fact(n int) int {

	r := 1

	for n > 1 {
		r = r * n
		n--
	}

	return r

}
