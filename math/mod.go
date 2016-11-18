package math

// Returns the floating point reminder of division operation.
func Mod(x, y float64) float64 {

	s := x < 0
	x = Abs(x)
	y = Abs(y)

	r := x - y*Floor(x/y)

	if s {
		return -r
	} else {
		return r
	}

}
