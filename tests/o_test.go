package tests

import (
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
