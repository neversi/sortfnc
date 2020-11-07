package sortfnc

// MergeSort has recursive behavior wherein the container is divided into two smaller parts
// and each of these parts are sorted explicitly. Eventually, those parts are merged
func MergeSort(s Sorter, left, right int) {
	if (left < right) {
		middle := left / 2 + (right) / 2	//Division index

		MergeSort(s, left, middle);		// Applying Mergesort for the left half of the container
		MergeSort(s, middle + 1, right);	// For the right half of the container
		compare(s, left, middle, right)		// Sorting and merging function
	}
	
}

// Sorts and "Merges" the container
func compare(s Sorter, left, middle, right int) {
	leftIndex, rightIndex := left, middle + 1

	for leftIndex <= middle && rightIndex <= right {
		if (s.Less(leftIndex, rightIndex) == true) {
			leftIndex++
		} else {
			index := rightIndex
			for index != leftIndex {
				s.Swap(index, index - 1)
				index--
			}

			leftIndex++
			middle++
			rightIndex++
		}
	}
}