package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodes := make([]*Node, 101)
	return rec(node, nodes)
}

func rec(node *Node, nodes []*Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{
		Val: node.Val,
	}
	nodes[node.Val] = newNode

	for _, n := range node.Neighbors {
		newN := nodes[n.Val]
		if newN == nil {
			newN = rec(n, nodes)
		}
		newNode.Neighbors = append(newNode.Neighbors, newN)
	}

	return newNode
}

func buildGraph(lists [][]int) *Node {
	nodes := make([]*Node, len(lists))
	for i := range lists {
		nodes[i] = &Node{
			Val: i + 1,
		}
	}
	for i, adjs := range lists {
		neighbors := make([]*Node, len(adjs))
		for j, n := range adjs {
			neighbors[j] = nodes[n-1]
		}
		nodes[i].Neighbors = neighbors
	}
	return nodes[0]
}

func main() {
	adjs := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	g := buildGraph(adjs)
	fmt.Print(cloneGraph(g))
}
