package sortfnc

// SelectionSort sorts the given container selecting the minimum element 
// and append to the start of the container in every iteration
func SelectionSort(s Sorter) {
	var unsortedIndex = 0 					// Container in the selection sort is divided into two parts -> sorted and unsorted
	for ; unsortedIndex < s.Len() - 1; unsortedIndex++ {	// Every iteration the size of sorted elements increases
		var iterMinIndex = unsortedIndex
		for i := unsortedIndex; i < s.Len(); i++ {
			if s.Less(i, iterMinIndex) == true {
				iterMinIndex = i		// Finding the minimum value in sub_container (unsortedIndex...s.Len()-1)
			}
		}
		s.Swap(iterMinIndex, unsortedIndex);		// Including final element to the sorted part of the container
		
	}
}
