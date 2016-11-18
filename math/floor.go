package math

// Returns the greatest integer value less than or equal to x.
func Floor(x float64) float64 {

	// special cases

	if x == 0 || x == NaN || x == InfPos || x == InfNeg {
		return x
	}

	// break x onto integer and fractional parts and take the corresponding the
	// integer one

	if x < 0 {
		int, frac := Modf(-x)
		if frac != 0.0 {
			int = int + 1
		}
		return -int
	}

	int, _ := Modf(x)
	return int

}
