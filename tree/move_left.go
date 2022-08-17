package tree

var nodeMatrix [][]*Node

// 先遍历一遍按层数构建一个数组
func traverse(node *Node, depth int) {
	if len(nodeMatrix) < depth {
		nodeMatrix = append(nodeMatrix, make([]*Node, 0))
	}

	nodeMatrix[depth-1] = append(nodeMatrix[depth-1], node)

	if node.Left != nil {
		traverse(node.Left, depth+1)
		node.Left = nil
	}

	if node.Right != nil {
		traverse(node.Right, depth+1)
		node.Right = nil
	}
}

// 然后从最后一层开始，按从左到右的顺序挨个把节点接到上一层节点上
func MoveLeft(root *Node) {
	traverse(root, 1)

	for i := len(nodeMatrix) - 1; i > 0; i-- {
		curLevelCount := len(nodeMatrix[i])
		for j := 0; j < curLevelCount; j++ {
			if j%2 == 0 {
				nodeMatrix[i-1][j/2].Left = nodeMatrix[i][j]
			} else {
				nodeMatrix[i-1][j/2].Right = nodeMatrix[i][j]
			}
		}
	}
}
