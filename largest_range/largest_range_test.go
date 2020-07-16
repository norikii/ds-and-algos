package largest_range

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var arrInput = []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
var arrExpected = []int{1, 7}

func TestLargestRange(t *testing.T) {
	res := LargestRange(arrInput)
	assert.Equal(t, res, arrExpected)
}

func TestLargestRange2(t *testing.T) {
	LRTT(arrInput)
}
