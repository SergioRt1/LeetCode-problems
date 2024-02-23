package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	hi, lo, mid, current := x, 0, 0, 0
	for lo < hi {
		mid = lo + (hi-lo)/2
		current = mid * mid
		if current == x {
			return mid
		} else if current > x { // take left
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	current = lo * lo
	if current > x {
		lo--
	}

	return lo
}

func main() {
	n := 5
	fmt.Println(mySqrt(n))
}
