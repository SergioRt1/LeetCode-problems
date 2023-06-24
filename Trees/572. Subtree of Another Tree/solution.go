package main

import "fmt"

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

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil {
		return false
	}
	if (root == nil && subRoot == nil) || (root != nil && subRoot == nil) {
		return true
	}
	if root.Val == subRoot.Val && areSame(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func areSame(a, b *TreeNode) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if a == nil && b == nil {
		return true
	}
	return a.Val == b.Val && areSame(a.Left, b.Left) && areSame(a.Right, b.Right)
}

func main() {
	root := makeTree([]int{3, 4, 5, 1, 2, null, null, null, null, 0}, 0)
	subRoot := makeTree([]int{4, 1, 2}, 0)

	fmt.Print(isSubtree(root, subRoot))
}
