package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, current := 0, 0, 0
	size := n + m
	// shift nums1 n positions
	copy(nums1[n:], nums1)
	i = n
	for i < size && j < n {
		if nums1[i] < nums2[j] {
			nums1[current] = nums1[i]
			i++
		} else {
			nums1[current] = nums2[j]
			j++
		}
		current++
	}
	for j < n {
		nums1[current] = nums2[j]
		j++
		current++
	}
}

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}

	fmt.Println(nums1)
	merge(nums1, 5, nums2, 1)
	fmt.Println(nums1)
}
