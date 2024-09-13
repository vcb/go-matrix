package matrix

import "testing"

func TestMatrix_Det(t *testing.T) {
	A := NewMatrixFromStr("1, 6, 8; 1, 3, 5; 8, 5, 4")
	det, err := A.Det()
	if err != nil {
		t.Error(err)
	}
	t.Log(det)
}

func TestMatrix_LU(t *testing.T) {
	A := NewMatrixFromStr("1, 6, 8; 1, 3, 5; 8, 5, 4")
	L, U, err := A.LU()
	if err != nil {
		t.Error(err)
	}
	t.Log(L)
	t.Log(U)

	B, err := L.Mul(U)
	if err != nil {
		t.Error(err)
	}
	if !Equals(A, B) {
		t.Error("A!=LU")
	}
	t.Log(B)
}
