package quick_sort

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var testDuplicatesArr = []int{10, 15, 7, 9, 10, 34, 38, 58, 30, 15, 34, 10}
var testDuplicatesArrSorted = []int{7, 9, 10, 10, 10, 15, 15, 30, 34, 34, 38, 58}
var testArrSorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var testArrUnSorted = []int{3, 2, 1, 5, 4, 6, 9, 8, 7}
var emptyArr []int
var oneElem = []int{1}
var twoElms = []int{2,-5}
var twoElmsSorted = []int{-5, 2}

func TestQuickSort(t *testing.T) {
	sortedArr := QuickSort(testArrUnSorted)
	assert.Equal(t, sortedArr, testArrSorted)
}

func TestQuickSort_WithDuplicates(t *testing.T) {
	sortedArr := QuickSort(testDuplicatesArr)
	assert.Equal(t, sortedArr, testDuplicatesArrSorted)
}

func TestQuickSort_WithEmptyArr(t *testing.T) {
	sortedArr := QuickSort(emptyArr)
	assert.Equal(t, sortedArr, emptyArr)
}

func TestQuickSort_WithOneItem(t *testing.T) {
	sortedArr := QuickSort(oneElem)
	assert.Equal(t, sortedArr, oneElem)
}

func TestQuickSort_TwoItems(t *testing.T) {
	sortedArr := QuickSort(twoElms)
	assert.Equal(t, sortedArr, twoElmsSorted)
}
