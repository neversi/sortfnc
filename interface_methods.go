package sortfnc

// Sorter is interface implemented by values which has an container characteristics and
// methods like Less, Len, Swap. The implementation of Sorter can call functions like
// Bubble, Insertion, Merge, Quick... *Sort which sorts the container by Less() (which is specified by the developer)
type Sorter interface {
	Less(i, j int) bool
	Len() int
	Swap(i, j int)
}
//MySliceInt and other data types are the slices, uncomment and implement them if you want to do that
type MySliceInt []int
// type MySliceFloat []float32
// type MySliceString []string

//Less compares two elements in the container
func (sliceInt MySliceInt) Less(i, j int) bool {
	return sliceInt[i] < sliceInt[j]
}
//Len returns the size of the container
func (sliceInt MySliceInt) Len() int {
	return len(sliceInt)
}
//Swap substitutes the elements of the container
func (sliceInt MySliceInt) Swap(i, j int) {
	temp := sliceInt[i]
	sliceInt[i] = sliceInt[j]
	sliceInt[j] = temp
}
