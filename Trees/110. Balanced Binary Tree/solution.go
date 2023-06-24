package main

import (
	"fmt"
)

const null = -1 << 31

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

func isBalanced(root *TreeNode) bool {
	_, response := DFS(root)
	return response
}

func DFS(root *TreeNode) (deep int, balanced bool) {
	if root == nil {
		return 0, true
	}
	lDeep, lResult := DFS(root.Left)
	rDeep, rResult := DFS(root.Right)

	return max(lDeep, rDeep) + 1, lResult && rResult && abs(lDeep-rDeep) <= 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := []int{1, 2, 2, 3, 3, null, null, 4, 4}
	tree := makeTree(root, 0)
	fmt.Print(isBalanced(tree))
}
