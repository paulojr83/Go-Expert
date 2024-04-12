package main

import (
	"fmt"
)

type Node struct {
	data        int32
	left, right *Node
}

func inOrderTraversal(root *Node, result *[]int32) {
	if root == nil {
		return
	}
	inOrderTraversal(root.left, result)
	*result = append(*result, root.data)
	inOrderTraversal(root.right, result)
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	var result [][]int32
	var queue []*Node
	root := &Node{data: 1}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		leftIndex := indexes[node.data-1][0]
		rightIndex := indexes[node.data-1][1]

		if leftIndex != -1 {
			node.left = &Node{data: leftIndex}
			queue = append(queue, node.left)
		}

		if rightIndex != -1 {
			node.right = &Node{data: rightIndex}
			queue = append(queue, node.right)
		}
	}

	for _, k := range queries {
		swapSubtrees(root, 1, k)
		var res []int32
		inOrderTraversal(root, &res)
		result = append(result, res)
	}

	return result
}

func swapSubtrees(root *Node, depth, k int32) {
	if root == nil {
		return
	}

	if depth%k == 0 {
		root.left, root.right = root.right, root.left
	}

	swapSubtrees(root.left, depth+1, k)
	swapSubtrees(root.right, depth+1, k)
}

func main() {
	// Sample input
	indexes := [][]int32{{2, 3}, {-1, -1}, {-1, -1}}
	queries := []int32{1}
	// Call the swapNodes function
	result := swapNodes(indexes, queries)
	// Print the result
	for _, res := range result {
		for _, val := range res {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
