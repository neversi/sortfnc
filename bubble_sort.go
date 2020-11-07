package sortfnc

// BubbleSort has n iterations through the container and in each iteration 
// compares two neighbour elements in container
func BubbleSort(s Sorter) {
	for i := 0; i < s.Len() - 1; i++ { 		// First loop
		for j := 0; j < s.Len() - 1; j++ { 	// Second loop
			if s.Less(j, j + 1) == false {
				s.Swap(j, j + 1)	// if arr[j] > arr[j + 1] -> swap values
			}
		}
	}
}