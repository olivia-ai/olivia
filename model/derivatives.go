package model

type Derivative struct {
	Delta matrix
	Adjustment matrix
}

// ComputeLastLayerDerivatives returns the derivatives of the last layer L
func (nn NN) ComputeLastLayerDerivatives(output matrix) Derivative {
	l := len(nn.Layers) - 1
	lastLayer := nn.Layers[l]

	// Compute derivative for the last layer of weights and biases
	cost := output.Difference(lastLayer)
	sigmoidDerivative := lastLayer.Multiplication(
		lastLayer.ApplyFunction(subtractsOne),
	)

	// Compute delta and the weights' adjustment
	delta := cost.ApplyFunction(multipliesByTwo).Multiplication(sigmoidDerivative)
	weights := nn.Layers[l-1].Transpose().DotProduct(delta)

	return Derivative{
		Delta:      delta,
		Adjustment: weights,
	}
}