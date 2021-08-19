package network

import "math"

// Sigmoid is the activation function
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// MultipliesByTwo takes a float and returns the float multiplied by two
func MultipliesByTwo(x float64) float64 {
	return 2 * x
}

// SubtractsOne takes a float and returns the float subtracted by one
func SubtractsOne(x float64) float64 {
	return x - 1
}
