package main

import "fmt"

const null = -1 << 31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Make DFS with a Stack to traverse tree
func levelOrder(node *TreeNode) [][]int {
	response := make([][]int, 0)
	if node != nil {
		levels := make(map[*TreeNode]int)
		q := make([]*TreeNode, 0)
		q = append(q, node)
		levels[node] = 0
		for len(q) > 0 {
			last := len(q) - 1
			elem := q[last]
			q = q[:last]
			currentLevel := levels[elem]
			if len(response) <= currentLevel {
				response = append(response, []int{})
			}
			response[currentLevel] = append(response[currentLevel], elem.Val)
			if elem.Right != nil {
				q = append(q, elem.Right)
				levels[elem.Right] = currentLevel + 1
			}
			if elem.Left != nil {
				q = append(q, elem.Left)
				levels[elem.Left] = currentLevel + 1
			}
		}
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
	fmt.Print(levelOrder(tree))
}
