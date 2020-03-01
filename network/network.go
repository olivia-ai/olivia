package network

import (
	"fmt"

	"github.com/gookit/color"
	"gopkg.in/cheggaaa/pb.v1"
)

type Network struct {
	Layers  []Matrix
	Weights []Matrix
	Biases  []Matrix
	Output  Matrix
	Rate    float64
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(rate float64, input, output Matrix, hiddensNodes ...int) Network {
	// Create the layers arrays and add the input values
	inputMatrix := input
	layers := []Matrix{inputMatrix}
	// Generate the hidden layer
	for _, hiddenNodes := range hiddensNodes {
		layers = append(layers, CreateMatrix(len(input), hiddenNodes))
	}
	// Add the output values to the layers arrays
	layers = append(layers, output)

	// Generate the weights and biases
	weightsNumber := len(layers) - 1
	var weights []Matrix
	var biases []Matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := Columns(layers[i]), Columns(layers[i+1])

		weights = append(weights, RandomMatrix(rows, columns))
		biases = append(biases, RandomMatrix(Rows(layers[i]), columns))
	}

	return Network{
		Layers:  layers,
		Weights: weights,
		Biases:  biases,
		Output:  output,
		Rate:    rate,
	}
}

// FeedForward executes forward propagation for the given inputs in the network
func (network *Network) FeedForward() Matrix {
	for i := 0; i < len(network.Layers)-1; i++ {
		layer, weights, biases := network.Layers[i], network.Weights[i], network.Biases[i]

		productMatrix := DotProduct(layer, weights)
		Sum(productMatrix, biases)
		ApplyFunction(productMatrix, Sigmoid)

		// Replace the output values
		network.Layers[i+1] = productMatrix
	}

	return network.Layers[len(network.Layers)-1]
}

func (network *Network) FeedForwardWithValue(input []float64) Matrix {
	network.Layers[0] = Matrix{input}
	return network.FeedForward()
}

// FeedBackward executes back propagation to adjust the weights for all the layers
func (network *Network) FeedBackward() {
	output := network.Output
	l := len(network.Layers) - 1
	lastLayer := network.Layers[l]

	// Compute derivative for the last layer of weights and biases
	error := Difference(output, lastLayer)
	sigmoidDerivative := Multiplication(
		lastLayer,
		ApplyFunction(lastLayer, SubstractOne),
	)

	delta := Multiplication(
		ApplyFunction(error, MultiplyByTwo),
		sigmoidDerivative,
	)
	weights := DotProduct(Transpose(network.Layers[l-1]), delta)

	l -= 1

	sigmoidDerivative2 := Multiplication(
		network.Layers[l],
		ApplyFunction(network.Layers[l], SubstractOne),
	)
	delta2 := Multiplication(
		DotProduct(
			delta,
			Transpose(network.Weights[l]),
		),
		sigmoidDerivative2,
	)

	weights2 := DotProduct(Transpose(network.Layers[l-1]), delta2)

	// Adjust the weights and biases
	network.Weights[l] = Sum(network.Weights[l], ApplyRate(weights, network.Rate))
	network.Biases[l] = Sum(network.Biases[l], ApplyRate(delta, network.Rate))
	network.Weights[l-1] = Sum(network.Weights[l-1], ApplyRate(weights2, network.Rate))
	network.Biases[l-1] = Sum(network.Biases[l-1], ApplyRate(delta2, network.Rate))
}

// ComputeError returns the average of all the errors after the training
func (network *Network) ComputeError() float64 {
	// Feed forward to compute the last layer's values
	network.FeedForward()
	lastLayer := network.Layers[len(network.Layers)-1]
	errors := Difference(network.Output, lastLayer)

	// Make the sum of all the errors
	var i int
	var sum float64
	for _, a := range errors {
		for _, e := range a {
			sum += e
			i++
		}
	}

	// Compute the average
	return sum / float64(i)
}

// Train trains the neural network with a given number of iterations by executing
// forward and back propagation
func (network *Network) Train(iterations int) {
	// Create the progress bar
	bar := pb.New(iterations).Postfix(fmt.Sprintf(
		" - %s",
		color.FgBlue.Render("Creating the neural network"),
	))
	bar.Format("(██░)")
	bar.SetMaxWidth(60)
	bar.ShowCounters = false
	bar.Start()

	// Train the network
	for i := 0; i < iterations; i++ {
		network.FeedForward()
		network.FeedBackward()

		// Increment the progress bar
		bar.Increment()
	}

	bar.Finish()
	// Print the error
	arrangedError := fmt.Sprintf("%.5f", network.ComputeError())
	fmt.Printf("The error rate is %s.\n", color.FgGreen.Render(arrangedError))
}
