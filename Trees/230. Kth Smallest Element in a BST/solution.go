package main

import "fmt"

const null = -1 << 31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	_, r := dfs(root, k, 0)
	return r
}

// dfs calculates for a node the k-th smallest node, nSmaller is the number of nodes that are smaller that are not part of the current subtree
// Returns the number of nodes below this node, and the Kth smallest value found
func dfs(node *TreeNode, k, nSmaller int) (int, int) {
	if node == nil {
		return 0, -1
	}
	belowLeft, leftResponse := dfs(node.Left, k, nSmaller)
	if leftResponse >= 0 {
		return 0, leftResponse
	}
	currentN := belowLeft + nSmaller + 1
	if currentN == k {
		return 0, node.Val
	}
	belowRight, rightResponse := dfs(node.Right, k, currentN)
	if rightResponse >= 0 {
		return 0, rightResponse
	}
	return belowLeft + belowRight + 1, -1
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
	tree := makeTree([]int{6, 3, 7, 2, 5, null, null, 1, null, 4, null}, 0)
	k := 5
	fmt.Print(kthSmallest(tree, k))
}
