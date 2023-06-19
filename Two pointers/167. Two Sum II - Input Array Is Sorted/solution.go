package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if s := numbers[i] + numbers[j]; s == target {
			return []int{i + 1, j + 1}
		} else if s > target {
			j--
		} else {
			i++
		}
	}
	return []int{1, 1}
}

func main() {
	numbers := []int{-1, 0}
	target := -1
	fmt.Print(twoSum(numbers, target))
}
