package network

import (
	"math/rand"
)

// Matrix is an alias for [][]float64
type Matrix [][]float64

// GenerateRandom return a generated matrix of dimensions rows and columns.
// The values contained in the matrix are random and contained between -1 and 1.
func GenerateRandom(rows, columns int) (matrix Matrix) {
	matrix = make(Matrix, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			matrix[i][j] = rand.Float64()*2.0 - 1.0
		}
	}

	return
}

// Generate returns an empty matrix of dimensions rows and columns.
func Generate(rows, columns int) (matrix Matrix) {
	matrix = make(Matrix, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
	}

	return
}

// Rows returns the number of rows of the given matrix.
func (matrix Matrix) Rows() int {
	return len(matrix)
}

// Columns returns the number of columns of the given matrix.
func (matrix Matrix) Columns() int {
	return len(matrix[0])
}

type indexedFunction = func(i, j int, x float64) float64
// ApplyIndexedFunction applies the given indexed function to the values of the matrix 
// and returns the result.
func (matrix Matrix) ApplyIndexedFunction(fn indexedFunction) Matrix {
	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			matrix[i][j] = fn(i, j, matrix[i][j])
		}
	}

	return matrix
}

// ApplyFunction applies a function to the values of the matrix and returns the resukt.
func (matrix Matrix) ApplyFunction(fn func(x float64) float64) Matrix {
	return matrix.ApplyIndexedFunction(func(i, j int, x float64) float64 {
		return fn(x)
	})
}

// ApplyRate returns a new matrix where all the values have been multiplied by the rate.
func (matrix Matrix) ApplyRate(rate float64) Matrix {
	return matrix.ApplyFunction(func(x float64) float64 {
		return rate * x
	})
}

// DotProduct processes the result of the dot product between matrix and matrix2 
// and returns it.
func (matrix Matrix) DotProduct(matrix2 Matrix) Matrix {
	if matrix.Columns() != matrix2.Rows() {
		panic("Cannot make dot product between these two matrix.")
	}

	resultMatrix := Generate(matrix.Rows(), matrix2.Columns())
	return resultMatrix.ApplyIndexedFunction(func(i, j int, x float64) float64 {
		var sum float64

		for k := 0; k < matrix.Columns(); k++ {
			sum += matrix[i][k] * matrix2[k][j]
		}

		return sum
	})
}

// Sum processes the sum between matrix and matrix2, which should be of the same 
// dimensions, and returns the result.
func (matrix Matrix) Sum(matrix2 Matrix) (resultMatrix Matrix) {
	ErrorNotSameSize(matrix, matrix2)

	resultMatrix = Generate(matrix.Rows(), matrix.Columns())
	return resultMatrix.ApplyIndexedFunction(func(i, j int, x float64) float64 {
		return matrix[i][j] + matrix2[i][j]
	})
}

// Difference processes and returns the difference between matrix and matrix2.
func (matrix Matrix) Difference(matrix2 Matrix) (resultMatrix Matrix) {
	ErrorNotSameSize(matrix, matrix2)

	resultMatrix = Generate(matrix.Rows(), matrix.Columns())
	return resultMatrix.ApplyIndexedFunction(func(i, j int, x float64) float64 {
		return matrix[i][j] - matrix2[i][j]
	})
}

// Multiplication processes and returns the multiplication between matrix and matrix2
func (matrix Matrix) Multiplication(matrix2 Matrix) (resultMatrix Matrix) {
	ErrorNotSameSize(matrix, matrix2)

	resultMatrix = Generate(matrix.Rows(), matrix.Columns())
	return resultMatrix.ApplyIndexedFunction(func(i, j int, x float64) float64 {
		return matrix[i][j] * matrix2[i][j]
	})
}

// Transpose returns the given matrix transposed
func (matrix Matrix) Transpose() (resultMatrix Matrix) {
	resultMatrix = Generate(matrix.Columns(), matrix.Rows())

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			resultMatrix[j][i] = matrix[i][j]
		}
	}

	return resultMatrix
}

// ErrorNotSameSize panics if the matrices do not have the same dimension
func ErrorNotSameSize(matrix, matrix2 Matrix) {
	if matrix.Rows() != matrix2.Rows() && matrix.Columns() != matrix2.Columns() {
		panic("These two matrices must have the same dimension.")
	}
}
