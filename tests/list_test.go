package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashiki/my-algorithms/common"
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

func TestInsertSortedList(t *testing.T) {
	node5 := list.Node{
		Value: 5,
	}

	node3 := list.Node{
		Value: 3,
	}

	node1 := list.Node{
		Value: 1,
	}

	node1.Next = &node3
	node3.Next = &node5
	node5.Next = &node1

	n := 0
	res1 := list.InsertSortedList(nil, n)
	assert.EqualValues(t, n, res1.Value)

	res2 := list.InsertSortedList(&node1, n)
	p := res2
	for {
		fmt.Printf("%d\n", p.Value)
		if res2.Value == p.Next.Value {
			break
		} else {
			p = p.Next
		}
	}
	assert.EqualValues(t, 0, res2.Value)
}

func TestReverseLinkedList(t *testing.T) {
	list1 := common.BuildLinkedList([]int{1, 2, 3, 4, 5})
	res1 := list.ReverseLinkedList(list1)
	common.PrintLinkedList(res1)

	list2 := common.BuildLinkedList([]int{1, 2, 3, 4, 5})
	res2 := list.ReverseLinkedList2(list2)
	common.PrintLinkedList(res2)
}

func TestStackList(t *testing.T) {
	l := &list.StackList{}
	l.Push(1)
	l.Push(2)
	r, ok := l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	l.Push(3)
	l.Push(4)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	l.Push(5)
	l.Push(6)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	l.Push(7)
	l.Push(8)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
	r, ok = l.Pop()
	fmt.Printf("%d,%t\n", r, ok)
}
