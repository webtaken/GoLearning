package compare

func Population(a byte) uint64 {
	sum := uint64(0)
	for a != 0 {
		a = a & (a - 1)
		sum++
	}
	return sum
}

func Compare2Bytes(a, b byte) uint64 {
	return Population(a ^ b)
}

func CompareSHA256(c1, c2 [32]byte) uint64 {
	sum := uint64(0)
	for i := 0; i < 32; i++ {
		sum += Compare2Bytes(c1[i], c2[i])
	}
	return sum
}
