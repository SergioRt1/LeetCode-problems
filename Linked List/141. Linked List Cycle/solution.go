package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type nothing struct{}

// floyd's algorithm
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func iterativeSolutionTraversal(head *ListNode) bool {
	visit := map[*ListNode]nothing{}
	for {
		if head == nil {
			return false
		}
		if _, seen := visit[head]; seen {
			return true
		}
		visit[head] = nothing{}
		head = head.Next
	}
}

// uses a recursive traversal of the list to find loop with a visit map
func dfsHelper(head *ListNode) bool {
	visit := map[*ListNode]nothing{}
	return dfs(head, visit)
}

func dfs(node *ListNode, visit map[*ListNode]nothing) bool {
	if node == nil {
		return false
	}
	if _, seen := visit[node]; seen {
		return true
	}
	visit[node] = nothing{}
	return dfs(node.Next, visit)
}

func makeList(list []int) *ListNode {
	head := &ListNode{}
	current := head
	for idx := 0; idx < len(list)-1; idx++ {
		current.Val = list[idx]
		current.Next = &ListNode{}
		current = current.Next
	}
	current.Val = list[len(list)-1]
	return head
}

func main() {
	list := makeList([]int{3, 2, 0, -4})
	fmt.Print(hasCycle(list))
}
