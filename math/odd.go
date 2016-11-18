package math

// Returns true if x is odd or false otherwise.
func Odd(x int) bool {
	return Mod(float64(x), 2) != 0
}
