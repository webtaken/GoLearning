package main

import "testing"

func TestMul1(t *testing.T) {
	correct := 11 * 5000000000
	result := Mul1(11, 5000000000)
	if result != correct {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, correct)
	}
}

func TestMul2(t *testing.T) {
	correct := 11 * 5000000000
	result := Mul2(11, 5000000000)
	if result != correct {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, correct)
	}
}
