package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func treeHeight(root *Node) int {
	if root == nil {
		return -1
	} else {
		leftHeight := treeHeight(root.left)
		rightHeight := treeHeight(root.right)

		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func main() {
	// Sample input tree
	root := &Node{
		data: 1,
		left: &Node{
			data: 2,
			left: &Node{
				data:  4,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &Node{
			data: 3,
			left: nil,
			right: &Node{
				data: 5,
				left: nil,
				right: &Node{
					data: 6,
					left: &Node{
						data:  7,
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
		},
	}

	fmt.Println("Height of the binary tree:", treeHeight(root)) // Output: 2
}
