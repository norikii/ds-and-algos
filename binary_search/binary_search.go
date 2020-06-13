package binary_search

// BinarySearch finds an item in ordered array by
// recursively checking if the item is smaller than
// the item located in the middle of the array
// and then returns true if item in array or
// false if item not in array
func BinarySearch(needle int, heyStack []int) bool {
	// checks if array is empty
	if len(heyStack) == 0 {
		return false
	}
	// finds the middle index
	middle := len(heyStack)/2
	// checks if the value of middle index is needle
	if heyStack[middle] == needle {
		return true
	}
	// checks if the value of middle index is less than needle
	// and then recursively look for the needle in shortened array
	if middle < needle {
		return BinarySearch(needle, heyStack[:middle])
	}
	// checks if the value of middle index is greater than needle
	// and then recursively look for the needle in shortened array
	if middle > needle {
		return BinarySearch(needle, heyStack[middle:])
	}

	return false
}