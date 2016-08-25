package sort

import (
	"math"

	"github.com/ivolynets/algorithms/data"
)

// Pre-generated Pratt sequence
var pratt = []int{
	1, 2, 3, 4, 6, 8, 9, 12, 16, 18, 24, 27, 32, 36, 48, 54, 64, 72, 81, 96,
	108, 128, 144, 162, 192, 216, 243, 256, 288, 324, 384, 432, 486, 512,
	576, 648, 729, 768, 864, 972, 1024, 1152, 1296, 1458, 1536, 1728, 1944,
	2048, 2187, 2304, 2592, 2916, 3072, 3456, 3888, 4096, 4374, 4608, 5184,
	5832, 6144, 6561, 6912, 7776, 8192, 8748, 9216, 10368, 11664, 12288, 13122,
	31104, 32768, 34992, 36864, 39366, 41472, 46656, 49152, 52488, 55296, 59049,
	62208, 65536, 69984, 73728, 78732, 82944, 93312,
}

// Sorts the array in ascending order using shell sort algorithm.
// -------------------------------------------
// Worst case performance:      O(n ^ 2)
// Best case performance:       O(n ^ (3 / 2))
// Average case performance:    O(n log n)
// -------------------------------------------
func ShellSort(a []data.Object) {
	ShellSortKnuth(a)
}

// Sorts the array in ascending order using shell sort algorithm with Shell's
// gap sequence. Shell sequence is calculated using formula floor(n / 2 ^ k),
// e.g. for n = 1000, we get the following sequence:
// [500, 250, 125, 62, 31, 15, 7, 3, 1].
func ShellSortShell(a []data.Object) {

	n := len(a)

	// calculate gap, the first element of the Shell sequence
	g := int(math.Floor(float64(n / 2)))

	for g > 0 {

		for i := g; i < n; i++ {

			for j := i; j >= g && a[j].Less(a[j-g]); j -= g {

				t := a[j]
				a[j] = a[j-g]
				a[j-g] = t

			}

		}

		// recalculate gap, next element from the Shell sequence
		g = int(math.Floor(float64(g / 2)))

	}

}

// Sorts the array in ascending order using shell sort algorithm with Pratt's
// gap sequence. Pratt sequence is successive numbers of the form
// 2 ^ p * 3 ^ q, e.g. [1, 2, 3, 4, 6, 8, 9, 12, ...].
func ShellSortPratt(a []data.Object) {

	n := len(a)

	// calculate gap, the first element of the Pratt sequence

	k := 0
	g := pratt[k]

	for g < n/3 {
		k++
		g = pratt[k]
	}

	for {

		for i := g; i < n; i++ {

			for j := i; j >= g && a[j].Less(a[j-g]); j -= g {

				t := a[j]
				a[j] = a[j-g]
				a[j-g] = t

			}

		}

		// recalculate gap, next element from the Pratt sequence

		if k > 0 {
			k--
			g = pratt[k]
		} else {
			return
		}

	}

}

// Sorts the array in ascending order using shell sort algorithm with Knuth's
// gap sequence. Knuth sequence is calculated using formula (3 ^ k - 1) / 2,
// e.g. [1, 4, 14, 40, 121, ...].
func ShellSortKnuth(a []data.Object) {

	n := len(a)

	// calculate gap, the first element of the Knuth sequence

	g := 1

	for g < n/3 {
		g = 3*g + 1
	}

	for g > 0 {

		for i := g; i < n; i++ {

			for j := i; j >= g && a[j].Less(a[j-g]); j -= g {

				t := a[j]
				a[j] = a[j-g]
				a[j-g] = t

			}

		}

		// recalculate gap, next element from the Knuth sequence
		g /= 3

	}

}
