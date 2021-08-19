package model

import (
	"github.com/olivia-ai/olivia/matrices"
)

// An alias for the matrix type
type matrix = matrices.Matrix

// NN contains the Layers, Weights, Biases of a neural network then the actual output values
// and the learning rate.
type NN struct {
	Layers  []matrix
	Weights []matrix
	Biases  []matrix
	Rate    float64
	Time    float64
}

// CreateRNN creates a new recurrent neural network by filling the matrixes 
// with the given sizes and returns it.
func CreateRNN(learningRate float64, inputLayers int, hiddenLayersNodes ...int) NN {
	layers := []matrix{
		{make([]float64, inputLayers)},
	}
	// Generate the hidden(s) layer(s) and add them to the layers slice
	for _, hiddenLayerNodes := range hiddenLayersNodes {
		layers = append(
			layers, 
			matrices.Generate(1, hiddenLayerNodes),
		)
	}
	// Add the output values to the layers slice
	layers = append(layers, make(matrix, inputLayers))

	// Generate the weights and biases
	weightsNumber := len(layers) - 1
	var weights []matrix
	var biases []matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := layers[i].Columns(), layers[i+1].Columns()

		weights = append(weights, matrices.Generate(rows, columns))
		biases = append(biases, matrices.GenerateRandom(layers[i].Rows(), columns))
	}

	return NN{
		Layers:  layers,
		Weights: weights,
		Biases:  biases,
		Rate:    learningRate,
	}
}

// FeedForward processes the forward propagation of the neural network considering
// that the first layer already contains the input.
func (nn *NN) FeedForward() {
	for i := 0; i < len(nn.Layers)-1; i++ {
		layer, weights, biases := nn.Layers[i], nn.Weights[i], nn.Biases[i]

		productMatrix := layer.DotProduct(weights)
		productMatrix.Sum(biases)
		productMatrix.ApplyFunction(sigmoid)

		// Replace the output values by the calculated ones
		nn.Layers[i+1] = productMatrix
	}
}