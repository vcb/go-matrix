package main

import "testing"

func TestLU(t *testing.T) {
	A := NewMatrixFromStr("1, 5, 8; 9, 55, 24; 4, 2, 0")
	if A == nil {
		t.Fatal("failed to initialize matrix")
	}

	L, U, err := A.LU()
	if err != nil {
		t.Fatal("failed to LU decompose A:", err)
	}

	LU, err := L.Mul(U)
	if err != nil {
		t.Fatal("failed to multiply matrices:", err)
	}
	if !Equals(A, LU) {
		t.Fatal("A does not equal LU:\n", LU)
	}
}