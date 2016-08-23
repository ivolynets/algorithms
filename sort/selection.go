package sort

import "github.com/ivolynets/algorithms/data"

// Sorts the array in ascending order using selection sort algorithm.
// -------------------------------------
// Worst case performance:      O(n ^ 2)
// Best case performance:       O(n ^ 2)
// Avarage case performance:    O(n ^ 2)
// -------------------------------------
func SelectionSort(a []data.Object) {

	for i, v := range a {

		m := i

		for j := i + 1; j < len(a); j++ {
			if a[j].Less(a[m]) {
				m = j
			}
		}

		a[i] = a[m]
		a[m] = v

	}

}
