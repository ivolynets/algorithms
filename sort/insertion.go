package sort

import "github.com/ivolynets/algorithms/data"

// Sorts the array in ascending order using insertion sort algorithm.
// -------------------------------------
// Worst case performance:      O(n ^ 2)
// Best case performance:       O(n)
// Avarage case performance:    O(n ^ 2 / 2)
// -------------------------------------
func InsertionSort(a []data.Object) {

	for i := 1; i < len(a); i++ {

		for j := i; j > 0 && a[j].Less(a[j-1]); j-- {

			t := a[j]
			a[j] = a[j-1]
			a[j-1] = t

		}

	}

}
