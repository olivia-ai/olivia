package network

// Derivative contains the derivatives of `z` and the adjustments
type Derivative struct {
	Delta      Matrix
	Adjustment Matrix
}

// ComputeLastLayerDerivatives returns the derivatives of the last layer L
func (network Network) ComputeLastLayerDerivatives() Derivative {
	l := len(network.Layers) - 1
	lastLayer := network.Layers[l]

	// Compute derivative for the last layer of weights and biases
	cost := Difference(network.Output, lastLayer)
	sigmoidDerivative := Multiplication(lastLayer, ApplyFunction(lastLayer, SubtractsOne))

	// Compute delta and the weights' adjustment
	delta := Multiplication(
		ApplyFunction(cost, MultipliesByTwo),
		sigmoidDerivative,
	)
	weights := DotProduct(Transpose(network.Layers[l-1]), delta)

	return Derivative{
		Delta:      delta,
		Adjustment: weights,
	}
}

// ComputeDerivatives returns the derivatives of a specific layer l defined by i
func (network Network) ComputeDerivatives(i int, derivatives []Derivative) Derivative {
	l := len(network.Layers) - 2 - i

	// Compute derivative for the layer of weights and biases
	delta := Multiplication(
		DotProduct(
			derivatives[i].Delta,
			Transpose(network.Weights[l]),
		),
		Multiplication(
			network.Layers[l],
			ApplyFunction(network.Layers[l], SubtractsOne),
		),
	)
	weights := DotProduct(Transpose(network.Layers[l-1]), delta)

	return Derivative{
		Delta:      delta,
		Adjustment: weights,
	}
}

// Adjust make the adjusts
func (network Network) Adjust(derivatives []Derivative) {
	for i, derivative := range derivatives {
		l := len(derivatives) - i

		network.Weights[l-1] = Sum(
			network.Weights[l-1],
			ApplyRate(derivative.Adjustment, network.Rate),
		)
		network.Biases[l-1] = Sum(
			network.Biases[l-1],
			ApplyRate(derivative.Delta, network.Rate),
		)
	}
}
