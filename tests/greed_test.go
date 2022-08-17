package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/greed"
)

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	assert.EqualValues(t, "9534330", greed.LargestNumber(nums))

	nums = []int{4, 42, 45, 56, 81, 76, 123}
	assert.EqualValues(t, "81765645442123", greed.LargestNumber(nums))
}
