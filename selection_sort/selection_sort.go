package selection_sort

// SelectionSort firstly finds the smallest element of the array
// then removes that element from array
// and lastly append that smallest value to the new array
// hence ordering the array in a simple but slow fashion
func SelectionSort(arr []int) []int {
	arrayLen := len(arr)
	var newArr []int
	for i := 0; i < arrayLen; i++ {
		smallest, index := findSmallest(arr)
		arr = removeSmallest(arr, index, smallest)
		newArr = append(newArr, smallest)
	}

	return newArr
}

// finds the smallest element in the array
// returns the value of the smallest element and its array
func findSmallest(arr []int) (int, int) {
	var smallest int
	var i int
	for index, item := range arr {
		if index == 0 {
			smallest = item
			i = index
		}
		if item < smallest {
			smallest = item
			i = index
		}
	}

	return smallest, i
}

// removes the smallest element from the array
// returns new shortened array
func removeSmallest(s []int, index int, smallest int) []int {
	// storing value of last index
	valueOfLastIndex := s[len(s)-1]
	// moving the smallest value to the last index
	s[len(s)-1] = smallest
	// moving the value of last index to the index of the smallest value
	s[index] = valueOfLastIndex
	// removing the last index which is the smallest value
	s = s[:len(s)-1]

	return s
}
