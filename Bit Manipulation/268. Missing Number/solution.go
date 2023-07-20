package main

import (
	"fmt"
	"sort"
)

type nothing struct{}

func missingNumber(nums []int) int {
	return xorImplementation(nums)
}

// Using a sum with xor, time O(n), space O(1)
func xorImplementation(nums []int) int {
	sum := len(nums)
	for i, v := range nums {
		sum ^= i ^ v
	}
	return sum
}

// Using a sum, time O(n), space O(1)
func sumImplementation(nums []int) int {
	sum := len(nums)
	for i, v := range nums {
		sum += i - v
	}
	return sum
}

// Using map collection as a set time O(n), space O(n)
func mapImplementation(nums []int) int {
	set := make(map[int]nothing, len(nums))
	for i := 0; i <= len(nums); i++ {
		set[i] = nothing{}
	}
	for _, v := range nums {
		delete(set, v)
	}
	for key := range set {
		return key
	}
	return 0
}

// Sorting time O(n lg n) space O(1)
func sortImp(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Print(missingNumber(nums))
}
