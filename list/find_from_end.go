package list

func FindFromEnd(head *Node, k int) *Node {
	fastP := head
	for i := 0; i < k; i++ {
		fastP = fastP.Next
	}

	slowP := head
	for fastP != nil {
		fastP = fastP.Next
		slowP = slowP.Next
	}

	return slowP
}
