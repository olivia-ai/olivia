package model

import (
	"math"
)

// sigmoid is the chosen activation function
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// SubtractsOne takes a float and returns the float subtracted by one
func subtractsOne(x float64) float64 {
	return 1 - x
}

func negativeLogLikelihood(expected, y []float64) float64 {
	for i, val := range expected {
		if val == 1 {
			return -math.Log(y[i])
		}
	}

	return 1.0
}

// softmax implementation https://en.wikipedia.org/wiki/Softmax_function
func softmax(x []float64) []float64 {
	var sum float64
	result := make([]float64, len(x))

	for i, v := range x {
		exp := math.Exp(v)
		result[i] = exp
		sum += exp
	}

	if sum != 0.0 {
		inverseSum := 1.0 / sum
		for i := range result {
			result[i] *= inverseSum
		}
	}

	return result
}