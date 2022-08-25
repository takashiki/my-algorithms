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

func TestBookMeeting(t *testing.T) {
	aSchedule := [][]string{{"9:30", "10:30"}, {"12:00", "13:00"}, {"16:00", "18:00"}}
	aWorkTime := []string{"9:00", "20:00"}
	bSchedule := [][]string{{"10:00", "11:32"}, {"12:30", "14:30"}, {"14:30", "15:00"}, {"16:00", "17:00"}}
	bWorkTime := []string{"10:00", "18:30"}
	minutes := 30

	available := o.BookMeeting(aSchedule, bSchedule, aWorkTime, bWorkTime, minutes)
	fmt.Printf("%#v\n", available)
	assert.EqualValues(t, [][]string{{"15:00", "16:00"}, {"18:00", "18:30"}}, available)
}

func TestRemoveIslands(t *testing.T) {
	input := [][]bool{
		{true, false, false, false, false, false},
		{false, true, false, true, true, true},
		{false, false, true, false, true, false},
		{true, true, false, false, true, false},
		{true, false, true, true, false, false},
		{true, false, false, false, false, true},
	}

	output := o.RemoveIslands(input)
	for _, row := range output {
		fmt.Printf("%#v\n", row)
	}
}

func TestLongestIdenticalList(t *testing.T) {
	s := "abcabcbb"
	expected := 3

	assert.EqualValues(t, expected, o.LongestIdenticalList(s))
}

func TestMaxGcd(t *testing.T) {
	nums := []int{789, 834, 980, 695, 1112}
	expected := 139

	assert.EqualValues(t, expected, o.MaxGcd(nums))
}
