package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/str"
)

func TestRevertString(t *testing.T) {
	in1 := "abcdefg"
	assert.EqualValues(t, "gfedcba", str.RevertString(in1))

	in2 := "我爱你"
	assert.EqualValues(t, "你爱我", str.RevertString(in2))
}
