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