package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/dp"
)

func TestFib(t *testing.T) {
	n := 11
	fibN := dp.Fib(n)
	assert.EqualValues(t, 89, fibN)
}

func TestMinDistance(t *testing.T) {
	s1 := "impact"
	s2 := "pac"

	minDistance := dp.MinDistance(s1, s2)
	assert.EqualValues(t, 3, minDistance)

	s1 = "apple"
	s2 = "banana"
	minDistance = dp.MinDistance(s1, s2)
	assert.EqualValues(t, 5, minDistance)
}

func TestLongestIncreasingList(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4

	assert.EqualValues(t, expected, dp.LongestIncreasingList(nums))
}

