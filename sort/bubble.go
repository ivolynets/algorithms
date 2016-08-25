package sort

import "github.com/ivolynets/algorithms/data"

// Sorts the array in ascending order using bubble sort algorithm.
// -------------------------------------
// Worst case performance:      O(n ^ 2)
// Best case performance:       O(n)
// Average case performance:    O(n ^ 2)
// -------------------------------------
func BubbleSort(a []data.Object) {

	n := len(a)

	for i := 0; i < n; i++ {

		for j := n - 1; j > i; j-- {

			if a[j].Less(a[j-1]) {

				t := a[j]
				a[j] = a[j-1]
				a[j-1] = t

			}

		}

	}

}
