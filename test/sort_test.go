package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/neversi/sortfnc"
)

func createRandomSlice(len int) []int {
	result := make([]int, 0, len)
	for i := 0; i < len; i++ {
		result = append(result, rand.Int() % 40)
	}
	return result
}
//TestSorted ...
func TestSorted(t *testing.T) {
	for i := 0; i < 15; i++ {
		result := createRandomSlice(10)
		myResult := make([]int, len(result))
		copy(myResult, []int(result))
		sort.Sort(sortfnc.MySliceInt(result))
		sortfnc.BubbleSort(sortfnc.MySliceInt(myResult))
	
		for i := 0; i < 10; i++ {
			if result[i] != myResult[i] {
				t.Error("Error in: ", result[i], "is not equal to ", myResult[i])
			}	
		}	
	}
}