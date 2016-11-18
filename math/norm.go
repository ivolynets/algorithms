package math

// Breaks float value x onto normalized fraction y and integral power of two e,
// so that x = y * 2^e.
func Norm(x float64) (y float64, e int) {

	// special cases

	if x == 0 || x == InfPos || x == InfNeg || x == NaN {
		return x, 0
	}

	const SmallestNormal = 2.2250738585072014e-308 // 2^(-1022)
	if Abs(x) < SmallestNormal {
		y, e = x*(1<<shift), -shift
	} else {
		y, e = x, 0
	}

	z := Float64bits(y)
	e += int((z>>shift)&mask) - bias + 1
	z &^= mask << shift
	z |= (-1 + bias) << shift
	y = BitsFloat64(z)

	return

}

// Reverse function of Norm. Takes normalized fraction x and integral power of
// two e and returns x * 2^e.
func Denorm(x float64, e int) float64 {

	// special cases

	if x == 0 || x == InfPos || x == InfNeg || x == NaN {
		return x
	}

	x, y := Norm(x)
	e += y

	f := Float64bits(x)
	e += int(f>>shift)&mask - bias

	if e < -1024 {
		return 0
	}

	if e > 1023 {
		if x < 0 {
			return InfNeg
		} else {
			return InfPos
		}
	}

	m := 1.0

	if e < -1022 {
		e += 52
		m = 1.0 / (1 << 52) // 2^(-52)
	}

	f &^= mask << shift
	f |= uint64(e+bias) << shift

	return m * BitsFloat64(f)

}
