package math

import "fmt"

// Calculates square root of x.
func Sqrt(x float64) (float64, error) {
	return SqrtNewton(x)
}

// Calculates square root of x using Newton's (Newton-Raphson) method. Newton's
// method is about approximating to the root by iterating over z values until
// needed precision is achieved. Each next z value is calculating using formula
//
//      z = z - f(z) / f'(z)
//
// where f(z) is some approximation function and f'(z) is derivative of this
// function. Initial z value is a guess which is reasonably close to the true
// root. Approximation function is
//
//      f(z) = z^2 - x
//
// where x is our original argument, therefore derivative is
//
//      f'(z) = 2*z
//
// When we are checking precision, we compare expected value calculated as
//
//      z - x / z
//
// with current z value multiplied by epsolon (desired precision).
func SqrtNewton(x float64) (float64, error) {

	// special cases

	switch {
	case x < 0:
		return NaN, fmt.Errorf("Argument cannot be negative")
	case x == 0:
		return 0, nil
	}

	// begin approximation

	e := 1e-15 // epsilon (desired precision)
	z := x

	if Abs(z-x/z) > e*z {
		z = z - (z*z-x)/(2*z)
	}

	return z, nil

}
