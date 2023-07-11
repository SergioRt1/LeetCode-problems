package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := recursive(head, n)
	if size == n {
		return head.Next
	}
	return head
}

// recursive returns the nth position of the current node from the end and remove the n value
func recursive(node *ListNode, n int) int {
	if node == nil {
		return 0
	}
	nth := recursive(node.Next, n) + 1
	if nth-1 == n {
		node.Next = node.Next.Next
	}
	return nth
}

func makeList(l []int, idx int) *ListNode {
	if idx >= len(l) {
		return nil
	}
	return &ListNode{
		Val:  l[idx],
		Next: makeList(l, idx+1),
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	n := 2
	head := makeList(list, 0)
	fmt.Print(removeNthFromEnd(head, n))
}
