package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nothing struct{}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre, in := 0, 0
	seen := make(map[int]nothing, len(preorder))

	var rec func() *TreeNode
	rec = func() *TreeNode {
		if pre >= len(preorder) {
			return nil
		}
		val := preorder[pre]
		pre++
		seen[val] = nothing{}
		node := &TreeNode{
			Val: val,
		}
		if val != inorder[in] {
			node.Left = rec()
		}
		in++ // advance in for middle node
		if in < len(inorder) {
			if _, ok := seen[inorder[in]]; !ok { // if is a new node
				node.Right = rec()
			}
		}

		return node
	}

	return rec()
}

func main() {
	pre := []int{3, 1, 2, 4}
	in := []int{1, 2, 3, 4}
	fmt.Println(pre, in)
	t := buildTree(pre, in)
	fmt.Println(t)
}
