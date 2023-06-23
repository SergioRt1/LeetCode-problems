package main

import (
	"fmt"
)

const null = -1 << 31

// TreeNode A definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(root []int, idx int) *TreeNode {
	if idx >= len(root) || root[idx] == null {
		return nil
	}
	return &TreeNode{
		Val:   root[idx],
		Left:  makeTree(root, idx*2+1),
		Right: makeTree(root, idx*2+2),
	}
}

// Use a DFS to walk throw the tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := []int{3, 9, 20, null, null, 15, 7}
	tree := makeTree(root, 0)
	fmt.Print(maxDepth(tree))
}
