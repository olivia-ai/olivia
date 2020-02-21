package network

import "math"

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidDerivative(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}

func Logit(x float64) float64 {
	return math.Log(x/1 - x)
}

func Twice(x float64) float64 {
	return 2 * x
}
