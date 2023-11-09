package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return binary(1, max, h, piles)
}

func binary(lo, hi, h int, piles []int) int {
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if validate(mid, h, piles) {
			result = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return result
}

func validate(k, h int, piles []int) bool {
	time := h
	for _, pile := range piles {
		time -= pile / k
		if pile%k != 0 {
			time -= 1
		}
		if time < 0 {
			return false
		}
	}
	return true
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	fmt.Println(minEatingSpeed(piles, h))
}
