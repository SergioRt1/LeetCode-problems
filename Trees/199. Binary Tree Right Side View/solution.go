package main

import "fmt"

const null = -1 << 32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use a modified BFS by level to determine the rightest element in each level
func rightSideView(root *TreeNode) []int {
	response := make([]int, 0)
	if root == nil {
		return response
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := make([]*TreeNode, 0)
		response = append(response, q[len(q)-1].Val)
		for _, node := range q {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		q = nextLevel
	}
	return response
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
	tree := makeTree([]int{}, 0)
	fmt.Print(rightSideView(tree))
}
