package network

import "math"

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
