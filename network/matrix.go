package network

import (
	"math/rand"
)

// RandomMatrix returns the value of a random matrix of *rows* and *columns* dimensions and
// where the values are between *lower* and *upper*.
func RandomMatrix(rows, columns int) (matrix [][]float64) {
	matrix = make([][]float64, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			matrix[i][j] = rand.Float64()*2.0 - 1.0
		}
	}

	return
}

// CreateMatrix returns an empty matrix which is the size of rows and columns
func CreateMatrix(rows, columns int) (matrix [][]float64) {
	matrix = make([][]float64, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns)
	}

	return
}

// Rows returns number of matrix's rows
func Rows(matrix [][]float64) int {
	return len(matrix)
}

// Columns returns number of matrix's columns
func Columns(matrix [][]float64) int {
	return len(matrix[0])
}

// ApplyFunctionWithIndex returns a matrix where fn has been applied with the indexes provided
func ApplyFunctionWithIndex(matrix [][]float64, fn func(i, j int, x float64) float64) [][]float64 {
	for i := 0; i < Rows(matrix); i++ {
		for j := 0; j < Columns(matrix); j++ {
			matrix[i][j] = fn(i, j, matrix[i][j])
		}
	}

	return matrix
}

// ApplyFunction returns a matrix where fn has been applied
func ApplyFunction(matrix [][]float64, fn func(x float64) float64) [][]float64 {
	return ApplyFunctionWithIndex(matrix, func(i, j int, x float64) float64 {
		return fn(x)
	})
}

// DotProduct returns a matrix which is the result of the dot product between matrix and matrix2
func DotProduct(matrix, matrix2 [][]float64) [][]float64 {
	if Columns(matrix) != Rows(matrix2) {
		panic("Cannot make dot product between these two matrix.")
	}

	return ApplyFunctionWithIndex(
		CreateMatrix(Rows(matrix), Columns(matrix2)),
		func(i, j int, x float64) float64 {
			var sum float64

			for k := 0; k < Columns(matrix); k++ {
				sum += matrix[i][k] * matrix2[k][j]
			}

			return sum
		},
	)
}

// Sum returns the sum of matrix and matrix2
func Sum(matrix, matrix2 [][]float64) [][]float64 {
	ErrorNotSameSize(matrix, matrix2)

	return ApplyFunctionWithIndex(matrix, func(i, j int, x float64) float64 {
		return matrix[i][j] + matrix2[i][j]
	})
}

// Difference returns the difference between matrix and matrix2
func Difference(matrix, matrix2 [][]float64) (resultMatrix [][]float64) {
	ErrorNotSameSize(matrix, matrix2)

	resultMatrix = CreateMatrix(Rows(matrix), Columns(matrix))

	return ApplyFunctionWithIndex(resultMatrix, func(i, j int, x float64) float64 {
		return matrix[i][j] - matrix2[i][j]
	})
}

// Multiplication returns the multiplication of matrix and matrix2
func Multiplication(matrix, matrix2 [][]float64) [][]float64 {
	ErrorNotSameSize(matrix, matrix2)

	return ApplyFunctionWithIndex(matrix, func(i, j int, x float64) float64 {
		return matrix[i][j] * matrix2[i][j]
	})
}

func Transpose(matrix [][]float64) (resultMatrix [][]float64) {
	resultMatrix = CreateMatrix(Columns(matrix), Rows(matrix))

	for i := 0; i < Rows(matrix); i++ {
		for j := 0; j < Columns(matrix); j++ {
			resultMatrix[j][i] = matrix[i][j]
		}
	}

	return resultMatrix
}

func ErrorNotSameSize(matrix, matrix2 [][]float64) {
	if Rows(matrix) != Rows(matrix2) && Columns(matrix) != Columns(matrix2) {
		panic("These two matrices must have the same dimension.")
	}
}
