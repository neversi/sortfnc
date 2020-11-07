package sortfnc

// QuickSort sorts the container comparing all elements to the "pivot" element
// and dividing them into two groups the left group which is the less than pivot
// and the right which the more or equal to the pivot element
func QuickSort(s Sorter, left, right int) {
	if (left < right) {
		pivot := pivotFinder(s, left, right)	// Sorting and selecting pivot element

		QuickSort(s, left, pivot - 1)		// Dividing the container in the range of left...pivot-1
		QuickSort(s, pivot + 1, right)		// Dividing the container in the range of pivot+1...right
	}
}

//Algorithm for selecting pivot element comparing to last indexed element and at meantime swap them (sort)
func pivotFinder(s Sorter, left, right int) int{
	pivotIndex := right

	i := left - 1

	for j := left; j <= right - 1; j++ {
		if s.Less(j, pivotIndex) {
			i++
			s.Swap(i, j)
		}
	}
	s.Swap(i + 1, pivotIndex)
	return (i + 1)
}