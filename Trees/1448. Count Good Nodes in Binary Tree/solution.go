package main

import "fmt"

const null = -1 << 31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return walk(root, -1<<31)
}

func walk(root *TreeNode, maxSeen int) int {
	if root == nil {
		return 0
	}
	var isGood int
	if root.Val >= maxSeen {
		isGood = 1
	}
	m := max(maxSeen, root.Val)
	return walk(root.Left, m) + walk(root.Right, m) + isGood
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func main() {
	tree := makeTree([]int{-1, 5, -2, 4, 4, 2, -2, null, null, -4, null, -2, 3, null, -2, 0, null, -1, null, -3, null, -4, -3, 3, null, null, null, null, null, null, null, 3, -3}, 0)
	fmt.Print(goodNodes(tree))
}
