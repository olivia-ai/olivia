package model

// computeLastLayerGradients returns the Gradients of the last layer L
func (nn NN) computeLastLayerGradients(output, truncatedOutput, exceptedOutput matrix) matrix {
	// Compute Gradient for the last layer of weights and biases
	// using the negative log likelihood loss function
	loss := exceptedOutput.Difference(output).ApplyRate(2)
	activationDerivative := output.ApplyFunction(reluDerivative)

	return activationDerivative.Multiplication(loss)
}

// ComputeGradients returns the gradients of a specific layer l defined by i
func (nn NN) ComputeGradients(i int, gradients []matrix, isLast bool) matrix {
	l := len(nn.Layers) - 2 - i

	weights := nn.Weights[l]
	// Link the encoder with the decoder by only passing by the second half of the decoder's
	// input, which represents the hidden state
	if isLast {
		weights = weights[len(weights)/2:]
	}

	// Compute Gradient for the layer of weights and biases
	return gradients[i].DotProduct(
		weights.Transpose(),
	).Multiplication(
		nn.Layers[l].Multiplication(
			nn.Layers[l].ApplyFunction(subtractsOne),
		),
	)
}

// Adjust takes the computed Gradients and applies the adjustments to
// the weights and biases
func (nn NN) Adjust(gradients []matrix) {
	for i, gradient := range gradients {
		l := len(gradients) - 2 - i
		if l < 0 {
			continue
		}

		nn.Weights[l] = nn.Weights[l].Sum(
			// Calculate weight adjustments
			nn.Layers[l].Transpose().DotProduct(gradient).ApplyRate(nn.Rate),
		)
		nn.Biases[l] = nn.Biases[l].Sum(
			gradient.ApplyRate(nn.Rate),
		)
	}
}