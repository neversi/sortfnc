package sortfnc

// ShakerSort as a BubbleSort sorts the container, however, it does it in both direction
func ShakerSort(s Sorter) {
	var leftSide, rightSide int = 0, s.Len()
	for i := 0; i < s.Len() / 2 + s.Len() % 2;  i++ {	// Left to Right sorting (the most right element is the biggest one)
		for j := leftSide; j < rightSide - 1; j++ {
			if s.Less(j, j + 1) == false {
				s.Swap(j, j + 1)
			}
		}
		rightSide-- 					// shrinking the range of sorting

		for j := rightSide -  1; j > leftSide; j-- {	// Right to Left sorting (the most left element is the lowest one)
			if s.Less(j, j - 1) != false {
				s.Swap(j, j - 1)
			}
		}
		leftSide++					// shrinking the range of sorting
	}
}
