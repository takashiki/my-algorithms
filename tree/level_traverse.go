package tree

import "fmt"

func LevelTraverse(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
