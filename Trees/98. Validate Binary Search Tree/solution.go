package main

import "fmt"

const inf = 1 << 32
const null = -inf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return DFS(root, -inf, inf)
}

func DFS(node *TreeNode, lower, maximum int) bool {
	if node == nil {
		return true
	}
	if node.Val > lower && node.Val < maximum {
		return DFS(node.Left, lower, node.Val) && DFS(node.Right, node.Val, maximum)
	}
	return false
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
	tree := makeTree([]int{5, 4, 6, null, null, 3, 7}, 0)
	fmt.Print(isValidBST(tree))
}
