package math

// Returns integer and fractional parts of float number so that
//
//      x = int + frac
//
// Float number IEEE mempry representation is:
//
//  ---------------------------------------------------------------------------
//   sign | exponent    | fraction
//  ---------------------------------------------------------------------------
//      X | XXXXXXXXXXX | XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//   bias | 01111111111 |
//  ---------------------------------------------------------------------------
//
// So,
//
//      exponent == 0  => float = -1^sign * 2^(-bias) * mantisa,
//      exponent > 0   => float = -1^sign * 2^(exponent - bias) * (1 + mantisa)
//
// where
//
//      mantisa  = (fraction / 2^shift)
//
func Modf(x float64) (int, frac float64) {

	// special cases

	if x < 1 {

		switch {
		case x < 0:
			int, frac = Modf(-x)
			return -int, -frac
		case x == 0:
			return x, x
		default:
			return 0, x
		}

	}

	b := Float64bits(x)             // take bits of argument
	e := uint(b>>shift)&mask - bias // extract exponent

	// keep the top 12+e bits, the integer part; clear the rest
	if e < shift {
		b &^= 1<<(shift-e) - 1
	}

	int = BitsFloat64(b) // restore integer part
	frac = x - int       // calculate fractional part

	return

}
