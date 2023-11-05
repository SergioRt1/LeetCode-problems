package main

import (
	"fmt"
	"math"
)

func isPalindromic(n, base int, digits []int) bool {
	idx := 0
	for n > 0 {
		digits[idx] = n % base
		n /= base
		idx++
	}
	for i, j := 0, idx-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	sqrtN := int(math.Sqrt(float64(n)))
	digits := make([]int, 40)
	some := false
	first := true

	for b := 2; b <= sqrtN; b++ {
		if isPalindromic(n, b, digits) {
			if first {
				fmt.Print(b)
				first = false
			} else {
				fmt.Print(" ", b)
			}

			some = true
		}
	}

	for k := n / (sqrtN + 1); k >= 1; k-- {
		b := (n - k) / k
		if b == sqrtN {
			continue
		}
		if (n-k)%k == 0 {
			if first {
				fmt.Print(b)
				first = false
			} else {
				fmt.Print(" ", b)
			}
			some = true
		}
	}
	if !some {
		fmt.Println("*")
	}
}
