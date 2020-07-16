package largest_range

import (
	"fmt"
	"sort"
)

// LargestRange get returns an array of length 2 representing the largest range of integers contained in the array
// the first number of the array should be the smallest number in the range and the second
// should be the highest number in the range
func LargestRange(arr []int) []int {
	var largestRangeArr []int
	var tempArr []int
	sort.Ints(arr)

	// if [i] == [i+1] + 1 then add to array
	for i:=0;i< len(arr);i++{
		if arr[i] != len(arr)-2 {
			if arr[i] == arr[i+1]+1 {
				tempArr = append(tempArr, arr[i])
			}
			if len(tempArr) > len(largestRangeArr) {
				largestRangeArr = tempArr
				tempArr = []int{}
			}
		}
	}

	return largestRangeArr
}

func LRTT(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr)
	for i:=0;i< len(arr);i++{
		if i < len(arr)-1 {
			//fmt.Println(arr[i])
			//fmt.Println(arr[i+1]+1)
			if arr[i]+1 == arr[i+1] {
				fmt.Println(arr[i+1])
			}
		}
	}
}




