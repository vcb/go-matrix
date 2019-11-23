package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Matrix is a basic matrix.
type Matrix struct {
	el         [][]*big.Float
	rows, cols int
}

// NewMatrix returns a m-by-n matrix.
func NewMatrix(rows, cols int) *Matrix {
	A := &Matrix{}
	A.el = make([][]*big.Float, rows)
	for i := range A.el {
		A.el[i] = make([]*big.Float, cols)
		for j := range A.el[i] {
			A.el[i][j] = new(big.Float)
		}
	}
	A.rows, A.cols = rows, cols
	return A
}

// NewMatrixFromStr returns a new matrix from a MATLAB-style string.
// Elements separated by commas, rows by semicolons, e.g. "1, 2; 3, 4".
func NewMatrixFromStr(s string) *Matrix {
	rows := strings.Split(s, ";")
	var els []string
	for _, row := range rows {
		els = append(els, strings.Split(row, ",")...)
	}
	if len(els) == 0 {
		return nil
	}

	A := NewMatrix(len(rows), len(els)/len(rows))
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			el := strings.TrimSpace(els[i*A.cols+j])
			x, err := strconv.ParseFloat(el, 64)
			if err != nil {
				return nil
			}
			A.Set(i, j, x)
		}
	}
	return A
}

// NewIdentityMatrix returns an identity matrix.
func NewIdentityMatrix(dim int) *Matrix {
	mI := NewMatrix(dim, dim)
	for ij := 0; ij < dim; ij++ {
		mI.Set(ij, ij, 1)
	}
	return mI
}

func (A *Matrix) String() string {
	s := []string{}
	for i := range A.el {
		l := []string{}
		for j := range A.el[i] {
			f, _ := A.el[i][j].Float64()
			l = append(l, fmt.Sprintf("% 7.3f", f))
		}
		s = append(s, fmt.Sprintf("%s",
			strings.Join(l, " ")))
	}
	return strings.Join(s, "\n")
}

// Copy returns a copy of the matrix.
func (A *Matrix) Copy() *Matrix {
	c := *A
	return &c
}

// Set sets the value at (i, j) to x.
func (A *Matrix) Set(i, j int, x float64) {
	A.el[i][j].SetFloat64(x)
}

// Mul returns the product A * B.
func (A *Matrix) Mul(B *Matrix) (*Matrix, error) {
	if A.cols != B.rows {
		return nil, fmt.Errorf("matrices have wrong dimensions")
	}

	C := NewMatrix(A.rows, B.cols)
	x := new(big.Float)
	for i := 0; i < A.rows; i++ {
		for k := 0; k < B.cols; k++ {
			for j := 0; j < A.cols; j++ {
				x.Mul(A.el[i][j], B.el[j][k])
				C.el[i][k].Add(C.el[i][k], x)
			}
		}
	}
	return C, nil
}

// Add returns the sum A + B.
func (A *Matrix) Add(B *Matrix) (*Matrix, error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, fmt.Errorf("matrices have different dimensions")
	}

	C := NewMatrix(A.rows, A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			C.el[i][j].Add(A.el[i][j], B.el[i][j])
		}
	}
	return C, nil
}

// Sub returns the result for A - B.
func (A *Matrix) Sub(B *Matrix) (*Matrix, error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, fmt.Errorf("matrices have different dimensions")
	}

	C := NewMatrix(A.rows, A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			C.el[i][j].Sub(A.el[i][j], B.el[i][j])
		}
	}
	return C, nil
}

// Rref returns a reduced row-echelon form of the matrix.
func (A *Matrix) Rref() *Matrix {
	B := A.Ref()

	q := new(big.Float)
	var i, j int
	for i < B.rows && j < B.cols {
		// column is zero, nothing to do
		if B.el[i][j].Cmp(big.NewFloat(0)) == 0 {
			j++
			continue
		}

		// divide row by the leading element
		q.Set(B.el[i][j])
		for k := j; k < B.cols; k++ {
			B.el[i][k].Quo(B.el[i][k], q)
		}

		// reduce all rows above
		if i > 0 {
			for h := i - 1; h >= 0; h-- {
				q.Quo(B.el[h][j], B.el[i][j])
				for k := j; k < B.cols; k++ {
					B.el[h][k].Sub(B.el[h][k], q)
				}
			}
		}

		i++
		j++
	}

	return B
}

// Ref returns a row-echelon form of the matrix using Gaussian elimination.
func (A *Matrix) Ref() *Matrix {
	B := A.Copy()

	q, x := new(big.Float), new(big.Float)
	var i, j int
	for i < B.rows && j < B.cols {
		// look for pivot row
		iMax := func(i int) int {
			var iMax int
			var max *big.Float
			for ; i < B.rows; i++ {
				b := new(big.Float).Copy(B.el[i][j])
				b.Abs(b)
				if max == nil || b.Cmp(max) == 1 {
					max, iMax = b, i
				}
			}
			return iMax
		}(i)
		if B.el[iMax][j].Cmp(big.NewFloat(0)) == 0 {
			j++
			continue
		}

		// swap rows
		if i != iMax {
			pivot := append([]*big.Float(nil), B.el[iMax]...)
			B.el[iMax] = B.el[i]
			B.el[i] = pivot
		}

		for h := i + 1; h < B.rows; h++ {
			q.Quo(B.el[h][j], B.el[i][j])
			B.Set(h, j, 0)
			for k := j + 1; k < B.cols; k++ {
				x.Mul(q, B.el[i][k])
				x.Sub(B.el[h][k], x)
				B.el[h][k].Copy(x)
			}
		}

		i++
		j++
	}
	return B
}

// Det calculates the determinant of the matrix.
func (A *Matrix) Det() (*big.Float, error) {
	if A.rows != A.cols {
		return nil, fmt.Errorf("matrix isn't square")
	}
	// TODO
	return new(big.Float), nil
}

// LU factors a square matrix as the product of lower and upper triangular matrices.
func (A *Matrix) LU() (*Matrix, *Matrix, error) {
	if A.rows != A.cols {
		return nil, nil, fmt.Errorf("matrix isn't square")
	}
	L, U := NewIdentityMatrix(A.rows), NewIdentityMatrix(A.rows)

	// based on http://lampx.tugraz.at/~hadley/num/ch2/2.3.php
	x := new(big.Float)
	for i := 0; i < A.rows; i++ {
		// upper
		for j := i; j < A.rows; j++ {
			U.el[i][j].Copy(A.el[i][j])

			for k := 0; k < i; k++ {
				x.Mul(L.el[i][k], U.el[k][j])
				U.el[i][j].Sub(U.el[i][j], x)
			}
		}

		// lower
		for j := i + 1; j < A.rows; j++ {
			L.el[j][i].Copy(A.el[j][i])

			for k := 0; k < i; k++ {
				x.Mul(L.el[j][k], U.el[k][i])
				L.el[j][i].Sub(L.el[j][i], x)
			}

			L.el[j][i].Quo(L.el[j][i], U.el[i][i])
		}
	}

	return L, U, nil
}

// Transpose returns the transpose of the matrix.
func (A *Matrix) Transpose() *Matrix {
	AT := NewMatrix(A.cols, A.rows)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			AT.el[j][i].Set(A.el[i][j])
		}
	}
	return AT
}

// Cat concatenates matrices A, B horizontally and returns [A|B]
func Cat(A, B *Matrix) *Matrix {
	if A.rows != B.rows {
		return nil
	}

	C := NewMatrix(A.rows, A.cols+B.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < C.cols; j++ {
			if j < A.cols {
				C.el[i][j].Set(A.el[i][j])
			} else {
				C.el[i][j].Set(B.el[i][j-A.cols])
			}
		}
	}
	return C
}

// Equals checks whether A = B.
func Equals(A, B *Matrix) bool {
	if A.rows != B.rows || A.cols != B.cols {
		return false
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if A.el[i][j].Cmp(B.el[i][j]) != 0 {
				return false
			}
		}
	}
	return true
}

// EstEquals checks whether A ~= B within eps.
func EstEquals(A, B *Matrix, eps float64) bool {
	if A.rows != B.rows || A.cols != B.cols {
		return false
	}

	diff, mean := new(big.Float), new(big.Float)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			diff.Sub(A.el[i][j], B.el[i][j])
			diff.Abs(diff)

			// use relative error if a and b are neither 0 nor infinity
			if A.el[i][j].MinPrec() != 0 && B.el[i][j].MinPrec() != 0 {
				mean.Add(A.el[i][j], B.el[i][j])
				mean.Quo(mean, big.NewFloat(2))
				diff.Quo(diff, mean)
			}

			f, _ := diff.Float64()
			fmt.Println(f)
			if f > eps {
				return false
			}
		}
	}
	return true
}
