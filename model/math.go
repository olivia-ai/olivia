package model

import (
	"math"
)

// sigmoid is the chosen activation function
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// MultipliesByTwo takes a float and returns the float multiplied by two
func multipliesByTwo(x float64) float64 {
	return 2 * x
}

// SubtractsOne takes a float and returns the float subtracted by one
func subtractsOne(x float64) float64 {
	return x - 1
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

// negativeLogLikelihood returns the NLL (negative log likelihood) loss function from the given
// values and expected values
func negativeLogLikelihood(values, expected []float64) float64 {
	if len(values) != len(expected) {
		panic("Values sets not the same length")
	}

	sum := 0.0
	for i := 0; i < len(values); i++ {
		sum += values[i] * expected[i]
	}

	return -math.Log(sum)
}