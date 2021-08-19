package model

import (
	"github.com/olivia-ai/olivia/matrix"
)

type Matrix = matrix.Matrix

// NeuralNetwork contains the Layers, Weights, Biases of a neural network then the actual output values
// and the learning rate.
type NeuralNetwork struct {
	Layers  []Matrix
	Weights []Matrix
	Biases  []Matrix
	Rate    float64
	Time    float64
}

// CreateNetwork creates a new NeuralNetwork by filling the matrixes with the given sizes and returns it.
func CreateNetwork(learningRate float64, inputLayers int, hiddensNodes ...int) NeuralNetwork {
	layers := []Matrix{
		{make([]float64, inputLayers)},
	}
	// Generate the hidden(s) layer(s) and add them to the layers slice
	for _, hiddenNodes := range hiddensNodes {
		layers = append(
			layers, 
			matrix.Generate(1, hiddenNodes),
		)
	}
	// Add the output values to the layers slice
	layers = append(layers, make(Matrix, inputLayers))

	// Generate the weights and biases
	weightsNumber := len(layers) - 1
	var weights []Matrix
	var biases []Matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := layers[i].Columns(), layers[i+1].Columns()

		weights = append(weights, matrix.Generate(rows, columns))
		biases = append(biases, matrix.GenerateRandom(layers[i].Rows(), columns))
	}

	return NeuralNetwork{
		Layers:  layers,
		Weights: weights,
		Biases:  biases,
		Rate:    learningRate,
	}
}

// FeedForward processes the forward propagation of the neural network considering
// that the first layer already contains the input.
func (nn *NeuralNetwork) FeedForward() {
	for i := 0; i < len(nn.Layers)-1; i++ {
		layer, weights, biases := nn.Layers[i], nn.Weights[i], nn.Biases[i]

		productMatrix := layer.DotProduct(weights)
		productMatrix.Sum(biases)
		productMatrix.ApplyFunction(Sigmoid)

		// Replace the output values by the calculated ones
		network.Layers[i+1] = productMatrix
	}
}