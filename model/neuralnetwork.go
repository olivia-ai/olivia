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
func CreateNetwork(learningRate float64, inputLayers int, outputLayers int, hiddensNodes ...int) NeuralNetwork {
	layers := []Matrix{
		{},
	}
	// Generate the hidden(s) layer(s) and add them to the layers slice
	for _, hiddenNodes := range hiddensNodes {
		layers = append(layers, matrix.Generate(inputLayers, hiddenNodes))
	}
	// Add the output values to the layers slice
	layers = append(layers, Matrix{})

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