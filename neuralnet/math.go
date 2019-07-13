package neuralnet

import "math"

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -float64(x)))
}

func SigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}
