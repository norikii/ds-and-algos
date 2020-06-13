package selection_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testArr = []int{10, 15, 7, 9, 10, 34, 38, 58, 30}
var testArrSorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var testArrUnSorted = []int{3, 2, 1, 5, 4, 6, 9, 8, 7}

func TestFindSmallest(t *testing.T) {
	smallest, _ := findSmallest(testArr)
	assert.Equal(t, 7, smallest)
}

func TestSelectionSort(t *testing.T) {
	sortedArr := SelectionSort(testArrUnSorted)
	assert.Equal(t, testArrSorted, sortedArr)
}


