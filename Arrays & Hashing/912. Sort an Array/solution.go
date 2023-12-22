package main

import "fmt"

func sortArray(nums []int) []int {
	second := make([]int, len(nums))
	copy(second, nums)
	return mergeSort(nums, second)
}

func mergeSort(nums []int, second []int) []int {
	if len(nums) <= 1 {
		return second
	}
	m := len(nums) >> 1

	left := mergeSort(second[m:], nums[m:])
	right := mergeSort(second[:m], nums[:m])

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			second[i+j] = left[i]
			i++
		} else {
			second[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		second[i+j] = left[i]
		i++
	}
	for j < len(right) {
		second[i+j] = right[j]
		j++
	}

	return second
}

func main() {
	nums := []int{5, 2, 3, 1, 8, 234, -123}
	fmt.Println(sortArray(nums))
}
