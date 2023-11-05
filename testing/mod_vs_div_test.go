package testing

import (
	"testing"
)

const (
	largeNumber = 100000000 // a large prime number
)

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeNumber / 1249999
	}
}

func BenchmarkModulus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeNumber % 1249999
	}
}
