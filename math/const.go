package math

// Mathematical constants

const (
	InfPos = float64(0x7FF0000000000000) // positive infinity representation
	InfNeg = float64(0xFFF0000000000000) // negative infinity representation

	NaN = float64(0x7FF8000000000001) // not-a-number constant

	// Euler's number is an important mathematical constant that is the base
	// of the natural logarithm. It can be calculated as the sum of the
	// infinite series
	//
	//     e = sum[n = 0..inf](1 / n!) = 1 + 1/1 + 1 / (1 * 2) + 1 / (1 * 2 * 3) + ...
	//
	E = 2.71828182845904523536028747135266249775724709369995957496696763

	Ln2 = 0.69314718055994530941723212145818
)

// Bit constants

const (
	shift = 52    // 52 bits shift
	mask  = 0x7FF // 00000000 00000000 00000000 00000000 00000000 00000000 00000111 11111111 (2047)
	bias  = 0x3FF // 00000000 00000000 00000000 00000000 00000000 00000000 00000011 11111111 (1023)
)
