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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//_, _, r := walk(root, p.Val, q.Val)
	//return r
	//return lowestCommonAncestorRecursive(root, p, q)
	return search(root, p, q)
}

// Linear solution with a full walk throw the tree, works for all trees
func walk(node *TreeNode, p, q int) (okP, okQ bool, r *TreeNode) {
	if node == nil {
		return false, false, nil
	}
	lp, lq, lr := walk(node.Left, p, q)
	rp, rq, rr := walk(node.Right, p, q)

	okP = lp || rp || node.Val == p
	okQ = lq || rq || node.Val == q

	if lr != nil {
		r = lr
	} else if rr != nil {
		r = rr
	} else if okP && okQ {
		r = node
	}

	return
}

// Logarithmic solution only works for binary search trees (sorted)
func search(root, p, q *TreeNode) *TreeNode {
	var small, big int
	if p.Val > q.Val {
		small, big = q.Val, p.Val
	} else {
		small, big = p.Val, q.Val
	}
	for root != nil {
		if root.Val > big {
			root = root.Left
		} else if root.Val < small {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}

// Recursive implementation
func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	var small, big int
	if p.Val > q.Val {
		small, big = q.Val, p.Val
	} else {
		small, big = p.Val, q.Val
	}
	return searchRecursive(root, small, big)
}

func searchRecursive(root *TreeNode, small, big int) *TreeNode {
	if root.Val > big {
		return searchRecursive(root.Left, small, big)
	} else if root.Val < small {
		return searchRecursive(root.Right, small, big)
	} else {
		return root
	}
}

func main() {
	root := makeTree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}, 0)
	p := &TreeNode{Val: 3}
	q := &TreeNode{Val: 9}
	fmt.Print(lowestCommonAncestor(root, p, q))
}
