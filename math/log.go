package math

// Returns natural logaritrhm
func Ln(x float64) float64 {
	return LnArtanh(x)
}

// Returns base-b logarithm of x.
func Log(b, x float64) float64 {
	return Ln(x) / Ln(b)
}

// Returns base-10 logarithm of x.
func Lg(x float64) float64 {
	return Log(10, x)
}

// Calculates natural logarithm using method based on the area hyperbolic
// tangent function. The formula is
//
//      ln(x) = 2 * sum[i = 0..inf](((x - 1) / (x + 1))^(2 * i + 1) / (2 * i + 1))
//
// We iterate until achieve maximum precision by testing each new value with
// previous one until they are equal.
//
// By default this method is not effective for large x, execution time linear
// depending on x, so we use identity
//
//      ln(xy) = ln(x) + ln(y)
//
// to improve performance. We break x onto normalized fraction y and integral
// power of two e, so that
//
//      x = y * 2^e
//
// Due to identity
//
//      ln(x^n) = n * ln(x)
//
// we have got a formula
//
//      ln(y * 2^e) = ln(y) + e * ln(2)
//
// where ln(2) is consat value. Finally we calculate logarithm only for small
// value y and add e * ln(2) to result.
func LnArtanh(x float64) float64 {

	// special cases

	switch {
	case x == 0:
		return InfNeg
	case x < 0:
		return NaN
	}

	f, e := Norm(x) // normalize argument x to f * 2^e

	// approximate log(f)

	r, t := 0.0, 0.0

	for i := 0.0; ; i++ {

		t = r + PowSqr((f-1)/(f+1), int(2*i+1))/(2*i+1)

		if t == r {
			break
		} else {
			r = t
		}

	}

	// scale back to obtain log(x)
	return 2*r + float64(e)*Ln2

}
