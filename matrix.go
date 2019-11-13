package main

import (
	"fmt"
	"math/big"
	"strings"
)

type Matrix struct {
	el         [][]*big.Float
	rows, cols int
}

// NewMatrix returns a m-by-n matrix
func NewMatrix(rows, cols int) *Matrix {
	m := &Matrix{}
	m.el = make([][]*big.Float, rows)
	for i := range m.el {
		m.el[i] = make([]*big.Float, cols)
		for j := range m.el[i] {
			m.el[i][j] = new(big.Float)
		}
	}
	m.rows, m.cols = rows, cols
	return m
}

func (m *Matrix) String() string {
	s := []string{}
	for i := range m.el {
		l := []string{}
		for j := range m.el[i] {
			l = append(l, fmt.Sprintf("% 3s", m.el[i][j].String()))
		}
		s = append(s, fmt.Sprintf("%s",
			strings.Join(l, " ")))
	}
	return strings.Join(s, "\n")
}

// Set sets the value at (i, j) to x
func (m *Matrix) Set(i, j int, x float64) {
	m.el[i][j].SetFloat64(x)
}

// Determinant calculates the determinant of a n-by-n matrix
// using LUP decomposition if possible
//
// det(A) = det(P^(-1)) * det(L) * det(U),
// where A = P^(-1) * L * U
func (m *Matrix) Determinant() (*big.Float, error) {
	if m.rows != m.cols {
		return nil, fmt.Errorf("matrix dimensions (%d, %d) are not equal", m.rows, m.cols)
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
	mT := NewMatrix(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			mT.el[j][i].Set(m.el[i][j])
		}
	}
	return mT
}

// Inverse returns the inverse of the matrix m
func (m *Matrix) Inverse() *Matrix {
	// TODO
	return &Matrix{}
}
