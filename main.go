package main

import (
	"fmt"
	"math/big"
	"strings"
)

// TODO TODO
// conditions for LU factorization https://arxiv.org/pdf/math/0506382v1.pdf

type Matrix struct {
	x [][]*big.Float
	n, m int
}

func main() {
	m := NewMatrix(3, 3)
	fmt.Println(m)
	fmt.Println(m.Determinant())
}

// NewMatrix returns a n-by-m matrix
func NewMatrix(dN, dM int) *Matrix {
	m := &Matrix{}
	m.x = make([][]*big.Float, dN)
	for i := range m.x {
		m.x[i] = make([]*big.Float, dM)
		for j := range m.x[i] {
			m.x[i][j] = new(big.Float)
		}
	}
	m.n, m.m = dN, dM
	return m
}

func (m *Matrix) String() string {
	s := []string{}
	for i := range m.x {
		l := []string{}
		for j := range m.x[i] {
			l = append(l, fmt.Sprintf("%s", m.x[i][j].String()))
		}
		s = append(s, fmt.Sprintf("| %s |",
			strings.Join(l, " ")))
	}
	return strings.Join(s, "\n")
}

// Determinant calculates the determinant of a n-by-n matrix
// using LUP decomposition if possible
// det(A) = det(P^(-1)) * det(L) * det(U),
// where A = P^(-1) * L * U
func (m *Matrix) Determinant() (*big.Float, error) {
	if m.m != m.n {
		return nil, fmt.Errorf("matrix dimensions (%d, %d) are not equal", m.m, m.n)
	}
	// l, u, pInv := m.LUPDecompose()
	// TODO
	return new(big.Float), nil
}

// LUPDecompose does stuff
func (m *Matrix) LUPDecompose() (*Matrix, *Matrix, *Matrix) {
	var l, u, pInv *Matrix
	// TODO
	return l, u, pInv
}

// Transpose returns the transposition of the matrix m
func (m *Matrix) Transpose() *Matrix {
	mT := NewMatrix(m.m, m.n)
	// TODO
	return mT
}

// Inverse returns the inverse of the matrix m
func (m *Matrix) Inverse() *Matrix {

}