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
	var resultRows [][]float64

	for row := 0; row < matrix.Rows(); row++ {
		var resultRow []float64

		for row2 := 0; row2 < matrix.Rows(); row2++ {
			var sum float64
			for column := 0; column < matrix.Columns(); column++ {
				sum += matrix.value[row][column] * matrix2.value[column][row2]
			}

			resultRow = append(resultRow, sum)
		}

		resultRows = append(resultRows, resultRow)
	}

	return Matrix{resultRows}
}
