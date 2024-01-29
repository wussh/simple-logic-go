package main

import (
	"fmt"
)

// TreeNode represents a node in the tree.
type TreeNode struct {
	Children []int
	Size     int32
}

// mostBalancedPartition finds the minimum difference between two partitions in a tree.
func mostBalancedPartition(parent []int32, files_size []int32) int32 {
	n := len(parent)
	tree := make([]TreeNode, n)

	// Build the tree
	for i := 0; i < n; i++ {
		if parent[i] != -1 {
			// Add child to the parent's Children list
			tree[parent[i]].Children = append(tree[parent[i]].Children, i)
		}
		// Set the size of the current node
		tree[i].Size = files_size[i]
	}

	var totalSize int32
	for _, size := range files_size {
		// Calculate the total size of all files
		totalSize += size
	}

	var minDiff int32 = 1<<31 - 1

	// Recursive function to perform depth-first search (DFS) on the tree
	var dfs func(node int) int32
	dfs = func(node int) int32 {
		// Initialize the subtree size with the size of the current node
		subtreeSize := tree[node].Size

		// Recursively calculate the subtree size by visiting each child
		for _, child := range tree[node].Children {
			subtreeSize += dfs(child)
		}

		// Calculate the difference between the total size and twice the subtree size
		diff := abs(totalSize - 2*subtreeSize)

		// Update the minimum difference if the current difference is smaller
		if diff < minDiff {
			minDiff = diff
		}

		// Return the total size of the subtree rooted at the current node
		return subtreeSize
	}

	// Start DFS from the root (node 0)
	dfs(0)

	// Return the minimum difference found
	return minDiff
}

// abs returns the absolute value of x.
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Sample input data
	parent := []int32{-1, 0, 1, 2}
	files_size := []int32{1, 2, 2, 1}

	// Find and print the result
	result := mostBalancedPartition(parent, files_size)
	fmt.Print(result)
}
