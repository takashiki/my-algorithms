package common

import "fmt"

func Min(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}

	return m
}

func Max(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}

func PrintLinkedList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Value)
		node = node.Next
	}

	fmt.Println("")
}

func BuildLinkedList(nums []int) *ListNode {
	cur := &ListNode{}
	var last *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		cur = &ListNode{
			Value: nums[i],
			Next:  last,
		}
		last = cur
	}

	return cur
}
