package network

import "math"

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func Twice(x float64) float64 {
	return 2 * x
}

func RemoveOne(x float64) float64 {
	return x - 1
}
