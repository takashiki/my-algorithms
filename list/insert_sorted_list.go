package list

// 1->3->5->1 输入成环有序链表，插入一个节点，使得链表仍然有序
func InsertSortedList(head *Node, n int) *Node {
	node := &Node{Value: n}
	// 当输入为 nil 时返回一个指向自身的节点
	if head == nil {
		node.Next = node
		return node
	}

	p := head
	for {
		// 遍历链表，当 p.val <= n <= p.next.val 时，插入到两者之间
		if p.Value <= n && n <= p.Next.Value {
			node.Next = p.Next
			p.Next = node
			return head
		}

		// 当遍历完之后如果都没有插入，那么说明 n 为最大值或最小值
		if p.Next.Value == head.Value {
			// 无论 n 是最大还是最小，都应指向 head
			p.Next = node
			node.Next = head
			// 当 n 为最小值时，node 为新的头节点
			if n < head.Value {
				return node
			} else {
				return head
			}
		}

		p = p.Next
	}
}
