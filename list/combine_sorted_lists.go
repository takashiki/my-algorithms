package list

type Node struct {
	Value int
	Next  *Node
}

func CombineSortedLists(a, b *Node) *Node {
	dummy := Node{}
	p := &dummy

	for {
		if a == nil || b == nil {
			break
		}

		if a.Value <= b.Value {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}

		p = p.Next
	}

	if a != nil {
		p.Next = a
	}

	if b != nil {
		p.Next = b
	}

	return dummy.Next
}
