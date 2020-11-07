package sortfnc

//InsertionSort sorting while inserting elements one by one
func InsertionSort(s Sorter) {
	for i := 1; i < s.Len(); i++ {
		for j := i - 1; j >= 0 && s.Less(i, j) == true; j-- {
			s.Swap(i, j);
		}
	}
}