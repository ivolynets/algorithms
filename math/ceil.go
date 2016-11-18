package math

// Returns the smallest integer value greater than or equal to x.
func Ceil(x float64) float64 {
	return -Floor(-x)
}
