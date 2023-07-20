package main

import (
	"fmt"
	"math/bits"
)

func reverseBits(num uint32) uint32 {
	return manuallyMove(num)
}

func lib(num uint32) uint32 {
	return bits.Reverse32(num)
}

func manually(num uint32) uint32 {
	newN := uint32(0)
	for i := 0; i < 32; i++ {
		newN |= (num >> i & 1) << (31 - i)
	}
	return newN
}

func manuallyMove(num uint32) uint32 {
	newN := uint32(0)
	for i := 0; i < 32; i++ {
		newN <<= 1
		newN |= num & 1
		num >>= 1
	}
	return newN
}

func printBinary(n uint32) {
	str := make([]byte, 32)
	for i := 0; i < 32; i++ {
		if n>>i&1 == 1 {
			str[31-i] = '1'
		} else {
			str[31-i] = '0'
		}
	}
	fmt.Println(string(str))
}

func main() {
	fmt.Print(reverseBits(43261596))
}
