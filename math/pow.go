package math

// Calculates x^y using exponentation by squaring method.
//
//      y is odd   => x^y = x * (x^2)^((y - 1) / 2)
//      y is even  => x^y = (x^2)^(y / 2)
//
func PowSqr(x float64, y int) float64 {

	// special cases

	switch {
	case y == 0 || x == 1:
		return 1
	case y < 0:
		return 1 / PowSqr(x, -y)
	}

	z := 1.0

	for y > 1 {
		if Odd(y) {
			z = x * z
			x = x * x
			y = (y - 1) / 2
		} else {
			x = x * x
			y = y / 2
		}
	}

	return x * z
}
