package neuralnet

import "math/rand"

// Returns the value of a matrix of *rows* and *columns* dimensions where all the
// values are equals to *value*.
func MakeMatrix(rows, columns int, value float64) [][]float64 {
	mat := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			mat[i][j] = value
		}
	}
	return mat
}

// Returns the value of a random matrix of *rows* and *columns* dimensions and
// where the values are between *lower* and *upper*.
func RandomMatrix(rows, columns int, lower, upper float64) [][]float64 {
	mat := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			mat[i][j] = rand.Float64()*(upper-lower) + lower
		}
	}
	return mat
}

// ApplyWeights returns the next layer where the weights have been applied to the values of the
// previous layer.
func ApplyWeights(layer []float64, weights [][]float64) (output []float64) {
	output = make([]float64, len(weights))

	for i := 0; i < len(weights); i++ {
		sum := 0.0
		for j := 0; j < len(layer); j++ {
			sum += layer[j] * weights[i][j]
		}
		output[i] = sum
	}

	return
}
