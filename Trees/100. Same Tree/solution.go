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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := makeTree([]int{1, 2, 1}, 0)
	q := makeTree([]int{1, 1, 2}, 0)

	fmt.Print(isSameTree(p, q))
}
