package popcount

import "testing"

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func BenchmarkPopCountv1(b *testing.B) {
	x := uint64(100)
	for i := 0; i < b.N; i++ {
		PopCountv1(x)
	}
}

func BenchmarkPopCountv2(b *testing.B) {
	x := uint64(100)
	for i := 0; i < b.N; i++ {
		PopCountv2(x)
	}
}

func BenchmarkPopCountv3(b *testing.B) {
	x := uint64(100)
	for i := 0; i < b.N; i++ {
		PopCountv3(x)
	}
}

func BenchmarkPopCountv4(b *testing.B) {
	x := uint64(100)
	for i := 0; i < b.N; i++ {
		PopCountv4(x)
	}
}
