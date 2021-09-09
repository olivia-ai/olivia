package model

type Gradient struct {
	Delta matrix
	Adjustment matrix
}

// ComputeLastLayerGradients returns the Gradients of the last layer L
func (nn NN) ComputeLastLayerGradients(output matrix) Gradient {
	l := len(nn.Layers) - 1
	lastLayer := nn.Layers[l]

	// Compute Gradient for the last layer of weights and biases
	// using the negative log likelihood loss function
	loss := negativeLogLikelihood(output[0], lastLayer[0])
	sigmoidGradient := lastLayer.Multiplication(
		lastLayer.ApplyFunction(subtractsOne),
	)

	// Compute delta and the weights' adjustment
	delta := sigmoidGradient.ApplyRate(loss)
	weights := nn.Layers[l-1].Transpose().DotProduct(delta)

	return Gradient{
		Delta:      delta,
		Adjustment: weights,
	}
}

// ComputeGradients returns the gradients of a specific layer l defined by i
func (nn NN) ComputeGradients(i int, gradients []Gradient) Gradient {
	l := len(nn.Layers) - 2 - i

	// Compute Gradient for the layer of weights and biases
	delta := gradients[i].Delta.DotProduct(
		nn.Weights[l].Transpose(),
	).Multiplication(
		nn.Layers[l].Multiplication(
			nn.Layers[l].ApplyFunction(subtractsOne),
		),
	)
	weights := nn.Layers[l-1].Transpose().DotProduct(delta)

	return Gradient{
		Delta:      delta,
		Adjustment: weights,
	}
}

// Adjust takes the computed Gradients and applies the adjustments to
// the weights and biases
func (nn NN) Adjust(gradients []Gradient) {
	for i, gradient := range gradients {
		l := len(gradients) - i

		nn.Weights[l-1] = nn.Weights[l-1].Sum(
			gradient.Adjustment.ApplyRate(nn.Rate),
		)
		nn.Biases[l-1] = nn.Biases[l-1].Sum(
			gradient.Delta.ApplyRate(nn.Rate),
		)
	}
}