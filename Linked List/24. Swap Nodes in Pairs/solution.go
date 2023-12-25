package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	head, head.Next, head.Next.Next = head.Next, head.Next.Next, head

	head.Next.Next = swapPairs(head.Next.Next)

	return head
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
	l := makeList([]int{}, 0)
	fmt.Println(swapPairs(l))
}
