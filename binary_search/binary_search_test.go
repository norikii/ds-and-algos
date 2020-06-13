package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testArr = []int{1, 5, 7, 9, 10, 34, 38, 58}

func TestBinarySearch_Present(t *testing.T) {
	isPresent := BinarySearch(7, testArr)
	assert.True(t, isPresent)
}

func TestBinarySearch_NotPresent(t *testing.T) {
	isPresent := BinarySearch(3, testArr)
	assert.False(t, isPresent)
}

func TestBinarySearch_EmptyArr(t *testing.T) {
	var emptyArr = []int{}
	isPreset := BinarySearch(1, emptyArr)
	assert.False(t, isPreset)
}

func TestBinarySearch_Zero(t *testing.T) {
	isPresent := BinarySearch(0, testArr)
	assert.False(t, isPresent)
}
