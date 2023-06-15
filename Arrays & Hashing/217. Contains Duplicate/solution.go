package main

import "fmt"

// map as set implementation
func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		} else {
			set[n] = struct{}{}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1}

	fmt.Print(containsDuplicate(nums))
}
