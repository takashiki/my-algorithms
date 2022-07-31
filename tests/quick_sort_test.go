package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/sort"
)

func TestQuickSort(t *testing.T) {
	values := []int{1, 6, 3, 5, 10}
	excepted := []int{1, 3, 5, 6, 10}

	sort.QuickSort(values, 0, len(values)-1)
	assert.EqualValues(t, excepted, values)
}
