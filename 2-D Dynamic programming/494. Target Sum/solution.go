package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	memo := make([]map[int]int, len(nums))
	for i := range memo {
		memo[i] = make(map[int]int, 0)
	}

	var rec func(i, v int) int
	rec = func(i, v int) int {
		if i >= len(nums) {
			if v == target {
				return 1
			} else {
				return 0
			}
		}
		if val, ok := memo[i][v]; ok {
			return val
		}

		currentNum := nums[i]
		value := rec(i+1, v-currentNum) + rec(i+1, v+currentNum)
		memo[i][v] = value

		return value
	}
	return rec(0, 0)
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	t := 3

	fmt.Println(findTargetSumWays(nums, t))
}
