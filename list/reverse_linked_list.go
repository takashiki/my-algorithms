package list

import (
	"github.com/takashiki/my-algorithms/common"
)

func ReverseLinkedList(head *common.ListNode) *common.ListNode {
	var prev, cur *common.ListNode
	cur = head

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func ReverseLinkedList2(head *common.ListNode) *common.ListNode {
	if head.Next == nil {
		return head
	}

	res := ReverseLinkedList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return res
}
