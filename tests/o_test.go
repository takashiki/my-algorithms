package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/o"
)

func TestMinimumDistance(t *testing.T) {
	needs := []string{"gym", "school", "store"}
	blocks := []o.Block{
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    true,
			"school": true,
			"store":  false,
		},
		{
			"gym":    true,
			"school": false,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  true,
		},
	}

	expectedIndex := 3
	foundIndex := o.FindTheBlock(blocks, needs)
	assert.EqualValues(t, expectedIndex, foundIndex)
}

func TestO1Stack(t *testing.T) {
	nums := []int{2, 3, 1, 4, 1}
	s := o.O1Stack{}
	for _, n := range nums {
		s.Push(n)
	}

	fmt.Printf("%#v\n", s)
	assert.EqualValues(t, 1, s.GetMin())

	d := s.Pop()
	assert.EqualValues(t, 1, d)
	assert.EqualValues(t, 1, s.GetMin())

	d = s.Pop()
	assert.EqualValues(t, 4, d)
	assert.EqualValues(t, 1, s.GetMin())

	d = s.Pop()
	assert.EqualValues(t, 1, d)
	assert.EqualValues(t, 2, s.GetMin())

	d = s.Pop()
	assert.EqualValues(t, 3, d)
	assert.EqualValues(t, 2, s.GetMin())
}
