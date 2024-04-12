package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func levelOrder(root *Node) {
	if root == nil {
		return
	}

	// Create a queue to store nodes
	queue := []*Node{root}

	for len(queue) > 0 {
		// Dequeue a node from the front of the queue
		node := queue[0]
		queue = queue[1:]

		// Print the data of the dequeued node
		fmt.Printf("%d ", node.data)

		// Enqueue the left child

		if node.left != nil {
			queue = append(queue, node.left)
		}

		// Enqueue the right child
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func main() {
	root := &Node{data: 1}
	root.right = &Node{data: 2}
	root.right.right = &Node{data: 5}
	root.right.right.left = &Node{data: 3}
	root.right.right.right = &Node{data: 6}
	root.right.right.left.right = &Node{data: 4}
	//root.right.right.left.left = &Node{data: 8}
	levelOrder(root)
}
