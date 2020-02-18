package network

import (
	"math/rand"
)

type Matrix struct {
	value [][]float64
}

// RandomMatrix returns the value of a random matrix of *rows* and *columns* dimensions and
// where the values are between *lower* and *upper*.
func RandomMatrix(rows, columns int) Matrix {
	mat := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			mat[i][j] = rand.Float64()*2.0 - 1.0
		}
	}

	return Matrix{mat}
}

func CreateMatrix(rows, columns int) Matrix {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
	}

	return Matrix{matrix}
}

func (matrix Matrix) Rows() int {
	return len(matrix.value)
}

func (matrix Matrix) Columns() int {
	return len(matrix.value[0])
}

func (matrix Matrix) DotProduct(matrix2 Matrix) Matrix {
	if matrix.Columns() != matrix2.Rows() {
		panic("Cannot make dot product between these two matrix.")
	}

	resultMatrix := CreateMatrix(matrix.Rows(), matrix2.Columns())

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix2.Columns(); j++ {
			var sum float64

			for k := 0; k < matrix.Columns(); k++ {
				sum += matrix.value[i][k] * matrix2.value[k][j]
			}

			resultMatrix.value[i][j] = sum
		}
	}

	return resultMatrix
}

func (matrix *Matrix) Add(matrix2 Matrix) {
	if matrix.Rows() != matrix2.Rows() && matrix.Columns() != matrix2.Columns() {
		panic("Cannot add these two matrix.")
	}

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			matrix.value[i][j] += matrix2.value[i][j]
		}
	}
}

func (matrix *Matrix) ApplyFunction(fn func(x float64) float64) {
	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			matrix.value[i][j] = fn(matrix.value[i][j])
		}
	}
}
