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

// CreateMatrix returns an empty matrix which is the size of rows and columns
func CreateMatrix(rows, columns int) Matrix {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
	}

	return Matrix{matrix}
}

// Rows returns number of matrix's rows
func (matrix Matrix) Rows() int {
	return len(matrix.value)
}

// Columns returns number of matrix's columns
func (matrix Matrix) Columns() int {
	return len(matrix.value[0])
}

// ApplyFunctionWithIndex returns a matrix where fn has been applied with the indexes provided
func (matrix Matrix) ApplyFunctionWithIndex(fn func(i, j int, x float64) float64) (resultMatrix Matrix) {
	resultMatrix = matrix

	for i := 0; i < resultMatrix.Rows(); i++ {
		for j := 0; j < resultMatrix.Columns(); j++ {
			resultMatrix.value[i][j] = fn(i, j, resultMatrix.value[i][j])
		}
	}

	return
}

// ApplyFunction returns a matrix where fn has been applied
func (matrix Matrix) ApplyFunction(fn func(x float64) float64) Matrix {
	return matrix.ApplyFunctionWithIndex(func(i, j int, x float64) float64 {
		return fn(x)
	})
}

// DotProduct returns a matrix which is the result of the dot product between matrix and matrix2
func (matrix Matrix) DotProduct(matrix2 Matrix) (resultMatrix Matrix) {
	if matrix.Columns() != matrix2.Rows() {
		panic("Cannot make dot product between these two matrix.")
	}

	resultMatrix = CreateMatrix(matrix.Rows(), matrix2.Columns()).
		ApplyFunctionWithIndex(func(i, j int, x float64) float64 {
			var sum float64

			for k := 0; k < matrix.Columns(); k++ {
				sum += matrix.value[i][k] * matrix2.value[k][j]
			}

			return sum
		})

	return
}

// Add returns the sum of matrix and matrix2
func (matrix Matrix) Add(matrix2 Matrix) Matrix {
	ErrorNotSameSize(matrix, matrix2)

	return matrix.ApplyFunctionWithIndex(func(i, j int, x float64) float64 {
		return matrix.value[i][j] + matrix2.value[i][j]
	})
}

// Substract returns the difference between matrix and matrix2
func (matrix Matrix) Substract(matrix2 Matrix) Matrix {
	ErrorNotSameSize(matrix, matrix2)

	return matrix.ApplyFunctionWithIndex(func(i, j int, x float64) float64 {
		return matrix.value[i][j] - matrix2.value[i][j]
	})
}

// Substract returns the multiplication of matrix and matrix2
func (matrix Matrix) Multiply(matrix2 Matrix) Matrix {
	ErrorNotSameSize(matrix, matrix2)

	return matrix.ApplyFunctionWithIndex(func(i, j int, x float64) float64 {
		return matrix.value[i][j] * matrix2.value[i][j]
	})
}

func (matrix Matrix) Transpose() (resultMatrix Matrix) {
	resultMatrix = CreateMatrix(matrix.Columns(), matrix.Rows())

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			resultMatrix.value[j][i] = matrix.value[i][j]
		}
	}

	return resultMatrix
}

func ErrorNotSameSize(matrix1, matrix2 Matrix) {
	if matrix1.Rows() != matrix2.Rows() && matrix2.Columns() != matrix2.Columns() {
		panic("These two matrices must have the same dimension.")
	}
}
