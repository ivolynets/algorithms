package math

// Calculates
func Exp(x float64) float64 {
	return ExpMaclaurin(x)
}

// Calculate the base-e exponential of x using Maclaurin series approximation.
// The formula is
//
//      exp(x) = sum[n = 0..inf](x^n / n!) => exp(x) = 1 + x + (x^2 / 2!) + (x^3 / 3!) + ...
//
// For more efficient calculation we reduce x first, so that
//
//      x = y + k * ln(2)
//
// where k is some positive integer. In this case we calculate only exp(y),
// where
//
//      y = x - k * ln(2)
//
// and then scale it back to exp(x) using reverese formula
//
//      exp(x) = exp(y) * 2^k
//
func ExpMaclaurin(x float64) float64 {

	const (
		Overflow  = 7.09782712893383973096e+02
		Underflow = -7.45133219101941108420e+02
		NearZero  = 1.0 / (1 << 28) // 1 / 2^(-28)
	)

	// special cases

	switch {
	case x == NaN || x == InfPos:
		return x
	case x == InfNeg:
		return 0
	case x > Overflow:
		return InfPos
	case x < Underflow:
		return 0
	case -NearZero < x && x < NearZero:
		return 1 + x
	}

	y, k := reduce(x) // reduce argument x

	// approximate exp(y)

	r, t := 0.0, 0.0

	for i := 0; ; i++ {

		t = r + PowSqr(y, i)/float64(Fact(i))

		if t == r {
			break
		} else {
			r = t
		}

	}

	// scale back to obtain exp(x)
	return scaleback(r, k)

}

// Reduces x so that x = y + k * ln(2). We assume y = 0.5 + ln(2)
func reduce(x float64) (y float64, k int) {

	const log2e = 1.4426950408889634073599246810019

	k = int(x*log2e - 0.5)
	y = x - float64(k)*Ln2

	return
}

// Scales exponent back using formula exp(x) = exp(y) * 2^k
func scaleback(e float64, k int) float64 {
	return e * PowSqr(2, k)
}
