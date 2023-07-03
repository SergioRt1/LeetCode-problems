package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return iterative(head)
}

// recursive solution
func recursive(head *ListNode) *ListNode {
	if head != nil {
		lastNode := recursive(head.Next)
		if head.Next != nil {
			head.Next.Next = head
		} else { // single node case
			lastNode = head
		}
		head.Next = nil
		return lastNode
	}
	return head
}

// Swap elements with a while loop
func iterative(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	var prev *ListNode
	for current.Next != nil {
		current.Next, prev, current = prev, current, current.Next
	}
	current.Next = prev
	return current
}

func buildList(list []int, idx int) *ListNode {
	if idx >= len(list) {
		return nil
	}
	return &ListNode{
		Val:  list[idx],
		Next: buildList(list, idx+1),
	}
}

func toList(head *ListNode) []int {
	response := make([]int, 0)
	for head != nil {
		response = append(response, head.Val)
		head = head.Next
	}
	return response
}

func main() {
	head := buildList([]int{1, 2, 3, 4, 5}, 0)
	result := reverseList(head)
	fmt.Print(toList(result))
}
