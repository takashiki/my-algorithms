package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/list"
)

func TestCombineSortedList(t *testing.T) {
	node5 := list.Node{
		Value: 5,
	}

	node4 := list.Node{
		Value: 4,
	}

	node3 := list.Node{
		Value: 3,
		Next:  &node4,
	}

	node1 := list.Node{
		Value: 1,
		Next:  &node3,
	}

	node2 := list.Node{
		Value: 2,
		Next:  &node5,
	}

	l := list.CombineSortedLists(&node1, &node2)
	var res [5]int
	i := 0
	for l != nil {
		res[i] = l.Value
		fmt.Printf("%d\n", l.Value)
		l = l.Next
		i++
	}

	assert.EqualValues(t, [5]int{1, 2, 3, 4, 5}, res)
}

func TestFindFromEnd(t *testing.T) {
	node5 := list.Node{
		Value: 5,
	}

	node4 := list.Node{
		Value: 4,
		Next:  &node5,
	}

	node3 := list.Node{
		Value: 3,
		Next:  &node4,
	}

	node2 := list.Node{
		Value: 2,
		Next:  &node3,
	}

	node1 := list.Node{
		Value: 1,
		Next:  &node2,
	}

	end2 := list.FindFromEnd(&node1, 2)
	assert.EqualValues(t, 4, end2.Value)
}
