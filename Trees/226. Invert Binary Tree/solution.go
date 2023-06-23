package main

import "fmt"

// TreeNode A definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(root []int, idx int) *TreeNode {
	if idx >= len(root) {
		return nil
	}
	return &TreeNode{
		Val:   root[idx],
		Left:  makeTree(root, idx*2+1),
		Right: makeTree(root, idx*2+2),
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}
}

func main() {
	root := []int{4, 2, 7, 1, 3, 6, 9}
	tree := makeTree(root, 0)
	fmt.Print(invertTree(tree))
}
