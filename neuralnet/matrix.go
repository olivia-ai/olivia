package neuralnet

import "math/rand"

// Returns the value of a matrix of *rows* and *columns* dimensions where all the
// values are equals to *value*.
func MakeMatrix(rows, colums int, value float64) [][]float64 {
	mat := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, colums)
		for j := 0; j < colums; j++ {
			mat[i][j] = value
		}
	}
	return mat
}

// Returns the value of a random matrix of *rows* and *columns* dimensions and
// where the values are between *lower* and *upper*.
func RandomMatrix(rows, colums int, lower, upper float64) [][]float64 {
	mat := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, colums)
		for j := 0; j < colums; j++ {
			mat[i][j] = rand.Float64()*(upper-lower) + lower
		}
	}
	return mat
}
