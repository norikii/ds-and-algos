package quick_sort

func QuickSort(unsorted []int) []int {
	var smallerThanPivot []int
	var greaterThanPivot []int
	var res []int

	// base case - if array contains only one element
	if len(unsorted) < 2 {
		return unsorted
	}

	// select the pivot for sorting the array
	pivot := unsorted[0]

	for _, item := range unsorted[1:] {
		if pivot < item {
			greaterThanPivot = append(greaterThanPivot, item)
		}
		if pivot >= item {
			smallerThanPivot = append(smallerThanPivot, item)
		}
	}

	//QuickSort(smallerThanPivot)
	//QuickSort(greaterThanPivot)

	res = append(res, QuickSort(smallerThanPivot)...)
	res = append(res, pivot)
	res = append(res, QuickSort(greaterThanPivot)...)

	// pick the element from the array - this will be the pivot
	// then find elements smaller than pivot and greater than pivot - this is called partitioning
	//

	return res
}


