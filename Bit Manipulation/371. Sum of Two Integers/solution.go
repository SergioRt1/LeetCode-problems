package main

import "fmt"

func getSum(a int, b int) int {
	var res, bA, bB, carry, sum int

	for i := 0; i < 64; i++ {
		bA = a >> i & 1
		bB = b >> i & 1
		sum = bA ^ bB ^ carry

		res |= sum << i

		carry = (bA & bB) | (carry & (bA ^ bB))
	}

	return res
}

func main() {
	fmt.Printf("%b\n", 14)
	fmt.Printf("%b\n", 7)
	fmt.Print(getSum(-2, 20))
}
