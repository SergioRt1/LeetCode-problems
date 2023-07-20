package main

import (
	"fmt"
	"math/bits"
)

func countBits(n int) []int {
	return memo(n)
}

// Use memorization to count with fewer operations O(n)
func memo(n int) []int {
	bitArray := make([]int, n+1)

	for i := 0; i <= n; i++ {
		bitArray[i] = bitArray[i>>1] + i&1
	}

	return bitArray
}

// Use internal lib bits to count O(n)
func lib(n int) []int {
	bitArray := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bitArray[i] = bits.OnesCount(uint(i))
	}
	return bitArray
}

// count bits manually O(n+bitSize), bitSize = 32 or 64 -> O(n)
func manually(n int) []int {
	bitArray := make([]int, n+1)
	for i := 0; i <= n; i++ {
		newN := i
		count := 0
		for newN != 0 {
			if newN&1 == 1 {
				count++
			}
			newN >>= 1
		}
		bitArray[i] = count
	}
	return bitArray
}

func main() {
	n := 5
	fmt.Print(countBits(n))
}
