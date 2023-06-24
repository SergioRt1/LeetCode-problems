package main

import "fmt"

const null = -1 << 31

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

func diameterOfBinaryTree(root *TreeNode) int {
	_, response := DFS(root)
	return response - 1
}

func DFS(root *TreeNode) (deepestPath, response int) {
	if root == nil {
		return 0, 0
	}
	lPath, lResponse := DFS(root.Left)
	rPath, rResponse := DFS(root.Right)

	maxResponse := max(rResponse, lResponse)

	return max(lPath, rPath) + 1, max(maxResponse, lPath+rPath+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := []int{1, 2, 3, 4, 5}
	tree := makeTree(root, 0)
	fmt.Print(diameterOfBinaryTree(tree))
}
