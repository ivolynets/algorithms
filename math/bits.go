package math

import "unsafe"

// Returns binary representation of float number.
func Float64bits(x float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&x))
}

// Returns float number from its binary representation.
func BitsFloat64(x uint64) float64 {
	return *(*float64)(unsafe.Pointer(&x))
}
