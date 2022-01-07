package model

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/olivia-ai/olivia/util"
)

func TestCreateNeuralNetwork(t *testing.T) {
	nn := CreateNeuralNetwork(0.5, 10, 20, 12, 30)

	for i := 0; i < 3; i++ {
		if nn.Layers[i].Columns() != nn.Weights[i].Rows() {
			t.Errorf("Error with the creation of the RNN, the size of layers/weights is wrong.")
		}
	}
}

func TestMNISTNN(t *testing.T) {
	nn := CreateNeuralNetwork(0.01, 784, 10, 100)

	// Retrieve the information from MNIST dataset and sort it
	// the input and output matrices
	var input, output matrix
	fmt.Println("retrieving the mnist data..")
	for _, lines := range util.ReadCSV("../data/mnist_train.csv") {
		number, _ := strconv.Atoi(lines[0])
		embedding := make([]float64, 10)
		embedding[number] = 1
		output = append(output, embedding)

		var image []float64
		for i := 1; i < 785; i++ {
			number, _ := strconv.ParseFloat(lines[i], 64)
			image = append(image, number)
		}
		input = append(input, image)
	}

	// Train the model and keep track of the process with a progress bar
	size := 10 * input.Rows()
	bar := util.ProgressBar("training", size)
	for epoch := 0; epoch < 10; epoch++ {
		for _, data := range input {
			y := nn.FeedForward(data)
			nn.PropagateBackward(
				nn.computeLastLayerGradients(y, negativeLogLikelihood(output[0], y[0])),
				false,
			)
			bar.Add(1)
		}
	}

	var loss float64
	testCSV := util.ReadCSV("../data/mnist_test.csv")
	for _, lines := range testCSV {
		number, _ := strconv.Atoi(lines[0])
		embedding := make([]float64, 10)
		embedding[number] = 1

		var image []float64
		for i := 1; i < 785; i++ {
			number, _ := strconv.ParseFloat(lines[i], 64)
			image = append(image, number)
		}

		actual := nn.FeedForward(image)
		loss += negativeLogLikelihood(embedding, actual[0])
	}
	fmt.Printf("\nloss: %f\n", loss/float64(len(testCSV)))
}
