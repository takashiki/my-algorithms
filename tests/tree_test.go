package tests

import (
	"fmt"
	"testing"

	"github.com/takashiki/my-algorithms/tree"
)

func TestMoveLeft(t *testing.T) {
	node6 := tree.Node{
		Value: 6,
	}

	node5 := tree.Node{
		Value: 5,
		Right: &node6,
	}

	node4 := tree.Node{
		Value: 4,
	}

	node3 := tree.Node{
		Value: 3,
		Right: &node5,
	}

	node2 := tree.Node{
		Value: 2,
		Right: &node4,
	}

	root := tree.Node{
		Value: 1,
		Left:  &node2,
		Right: &node3,
	}
	
	tree.MoveLeft(&root)

	traverse(&root)
}

func traverse(root *tree.Node) {
	if root.Left != nil {
		fmt.Printf("%d left value %d\n", root.Value, root.Left.Value)
		traverse(root.Left)
	}

	if root.Right != nil {
		fmt.Printf("%d right value %d\n", root.Value, root.Right.Value)
		traverse(root.Right)
	}
}
