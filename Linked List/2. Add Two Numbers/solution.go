package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rem int
	var current *ListNode

	head := &ListNode{}
	prev := head

	for l1 != nil && l2 != nil {
		current = &ListNode{}
		sum := l1.Val + l2.Val + rem
		current.Val = sum % 10
		rem = sum / 10
		prev.Next = current

		l1 = l1.Next
		l2 = l2.Next
		prev = current
	}

	for l1 != nil {
		current = &ListNode{}
		sum := l1.Val + rem
		current.Val = sum % 10
		rem = sum / 10
		prev.Next = current

		l1 = l1.Next
		prev = current
	}

	for l2 != nil {
		current = &ListNode{}
		sum := l2.Val + rem
		current.Val = sum % 10
		rem = sum / 10
		prev.Next = current

		l2 = l2.Next
		prev = current
	}

	if rem != 0 {
		current = &ListNode{Val: rem}
		prev.Next = current
	}

	return head.Next
}

func main() {
	list1 := createList([]int{9, 9, 9, 9, 9, 9, 9}, 0)
	list2 := createList([]int{9, 9, 9, 9}, 0)
	printList(list1)
	printList(list2)

	result := addTwoNumbers(list1, list2)

	printList(result)
}

func createList(list []int, idx int) *ListNode {
	if idx >= len(list) {
		return nil
	}
	return &ListNode{
		Val:  list[idx],
		Next: createList(list, idx+1),
	}
}

func printList(node *ListNode) {
	fmt.Print("[")
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(", ")
		}
		node = node.Next
	}
	fmt.Println("]")
}
