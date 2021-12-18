package model

import (
	"math/rand"
	"testing"
)

func TestCreateNeuralNetwork(t *testing.T) {
	nn := CreateNeuralNetwork(0.5, 10, 20, 12, 30)

	for i := 0; i < 3; i++ {
		if nn.Layers[i].Columns() != nn.Weights[i].Rows() {
			t.Errorf("Error with the creation of the RNN, the size of layers/weights is wrong.")
		}
	}
}

func TestNNFeedForward(t *testing.T) {
	nn := CreateNeuralNetwork(0.01, 3, 1)
	rand.Seed(5)

	input := matrix{
		{0, 1, 1},
		{0, 1, 0},
		{1, 0, 1},
	}
	output := matrix{
		{1},
		{0},
		{1},
	} 

	for epoch := 0; epoch < 100; epoch++ {
		for _, data := range input {
			y := nn.FeedForward(data)
			nn.PropagateBackward(nn.computeLastLayerGradients(y, negativeLogLikelihood(output[0], y[0])), false)
		}
	}
}