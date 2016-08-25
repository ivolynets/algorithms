package search

import "github.com/ivolynets/algorithms/data"

// Determines and retuens index of an element in the array using binary search
// algorithm. Returns -1 if element has not been found.
// -----------------------------------
// Worst case performance:    O(log n)
// Best case performance:     O(1)
// Average case performance:  O(log n)
// -----------------------------------
func BinarySearch(a []data.Object, i data.Object) int {

	l, h := 0, len(a)-1

	for l <= h {

		r := l + (h-l)/2

		switch c := i.Compare(a[r]); {
		case c < 0:
			h = r - 1
		case c > 0:
			l = r + 1
		default:
			return r
		}

	}

	return -1

}
