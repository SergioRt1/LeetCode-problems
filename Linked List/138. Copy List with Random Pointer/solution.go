package main

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	newHead, nodeMap := createList(head, 0)
	newCurrent := newHead
	current := head
	for current != nil {
		newCurrent.Random = nodeMap[current.Random]

		current = current.Next
		newCurrent = newCurrent.Next
	}

	return newHead
}

func createList(node *Node, idx int) (*Node, map[*Node]*Node) {
	if node == nil {
		return nil, make(map[*Node]*Node, idx)
	}

	child, nodes := createList(node.Next, idx+1)

	newNode := &Node{
		Val:  node.Val,
		Next: child,
	}
	nodes[node] = newNode

	return newNode, nodes
}
