package math

// Returns absolute value of x.
func Abs(x float64) float64 {

	if x <= 0 {
		return float64(0) - x
	} else {
		return x
	}

}
