package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode

	//Select the first element
	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head = list1
			current = head
			list1 = list1.Next
		} else {
			head = list2
			current = head
			list2 = list2.Next
		}
	} else if list1 != nil {
		head = list1
		current = head
		list1 = list1.Next
	} else if list2 != nil {
		head = list2
		current = head
		list2 = list2.Next
	}

	// Merge
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Clear reminding
	for list1 != nil {
		current.Next = list1
		list1 = list1.Next
		current = current.Next
	}
	for list2 != nil {
		current.Next = list2
		list2 = list2.Next
		current = current.Next
	}

	return head
}

func makeList(list []int, idx int) *ListNode {
	if idx >= len(list) {
		return nil
	}
	return &ListNode{
		Val:  list[idx],
		Next: makeList(list, idx+1),
	}
}

func main() {
	list1 := makeList([]int{1, 2, 4}, 0)
	list2 := makeList([]int{1, 3, 4}, 0)

	fmt.Print(mergeTwoLists(list1, list2))

	fmt.Print(mergeTwoLists(nil, nil))

	list3 := makeList([]int{1}, 0)
	fmt.Print(mergeTwoLists(nil, list3))
}
