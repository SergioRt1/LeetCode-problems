package main

import "fmt"

func maxSubArray(nums []int) int {
	return divideAndConquer(nums)
}

// O(n)
func kadaneAlgo(nums []int) int {
	best := nums[0]
	sum := 0
	for _, e := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += e
		if best < sum {
			best = sum
		}
	}
	return best
}

// O(n log(n))
func divideAndConquer(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	mid := len(nums) / 2
	l := divideAndConquer(nums[:mid])
	r := divideAndConquer(nums[mid:])
	m := findMaxMid(nums, mid) // O(n)

	return max(l, max(r, m))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const Inf = 1 << 30

func findMaxMid(nums []int, mid int) int {
	sum := 0
	bestR := -Inf
	for i := mid + 1; i < len(nums); i++ {
		sum += nums[i]
		if bestR < sum {
			bestR = sum
		}
	}
	sum = 0
	bestL := -Inf
	for i := mid; i >= 0; i-- {
		sum += nums[i]
		if bestL < sum {
			bestL = sum
		}
	}

	return max(bestL, max(bestR, bestL+bestR))
}

func main() {
	nums := []int{-2, -1}
	fmt.Println(maxSubArray(nums))
}
