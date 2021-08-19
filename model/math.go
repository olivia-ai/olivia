package model

import "math"

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
