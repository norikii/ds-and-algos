package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 0, 1, 2}
 	fmt.Println(findMaxNumLoop(arr))
	//tttt(arr)
}

func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sum(arr[1:])
}

func itemsInList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return 1 + itemsInList(arr[1:])
}

func findMaxNum(arr []int) int {
	max := arr[0]

	// base case
	if len(arr) == 0 {
		return max
	}

	if arr[0] < max {
		arr = arr[1:]
	}

	if arr[0] > max {
		max = arr[0]
		arr = arr[1:]
	}

	return findMaxNum(arr)
}

func findMaxNumLoop(arr []int) int {
	max := 0
	arrLen := len(arr)

	for i := 0; i< arrLen; i++  {
		// base case
		//if len(arr) == 0 {
		//	return max
		//}

		if arr[0] < max {
			arr = arr[1:]
		}

		if arr[0] > max {
			fmt.Println(arr[0])
			max = arr[0]
			arr = arr[1:]
		}
	}

	return max
}