package network

type Network struct {
	Layers  []Matrix
	Weights []Matrix
	Biases  []Matrix
	Output  Matrix
	Rate    float64
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(rate float64, input, output Matrix, hiddenNodes int) Network {
	// Create the layers arrays and add the input values
	inputMatrix := input
	layers := []Matrix{inputMatrix}
	// Generate the hidden layer
	layers = append(layers, CreateMatrix(len(input), hiddenNodes))
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
func (network *Network) FeedForward() {
	for i := 0; i < len(network.Layers)-1; i++ {
		layer, weights, biases := network.Layers[i], network.Weights[i], network.Biases[i]

		productMatrix := DotProduct(layer, weights)
		Sum(productMatrix, biases)

		// Replace the input values
		network.Layers[i+1] = productMatrix
		ApplyFunction(productMatrix, Sigmoid)

		// Replace the output values
		network.Layers[i+1] = productMatrix
	}
}

// FeedBackward executes back propagation to adjust the weights for all the layers
func (network *Network) FeedBackward() {
	output := network.Output

	// Compute derivative for the last layer of weights and biases
	lastLayer := network.Layers[2]
	error := Difference(output, lastLayer)
	sigmoidDerivative := Multiplication(lastLayer, ApplyFunction(lastLayer, SubstractOne))

	z := Multiplication(
		ApplyFunction(error, MultiplyByTwo),
		sigmoidDerivative,
	)
	weights := DotProduct(Transpose(network.Layers[1]), z)

	// Adjust the weights and biases
	network.Weights[1] = Sum(network.Weights[1], weights)
	network.Biases[1] = Sum(network.Biases[1], z)

	// Compute derivative for the first layer of weights and biases
	z = Multiplication(
		DotProduct(
			z,
			Transpose(network.Weights[1]),
		),
		Multiplication(
			network.Layers[1],
			ApplyFunction(network.Layers[1], SubstractOne),
		),
	)

	// Adjust the weights and biases
	network.Weights[0] = Sum(
		network.Weights[0],
		DotProduct(
			Transpose(network.Layers[0]),
			z,
		),
	)
	network.Biases[0] = Sum(network.Biases[0], z)
}

// Train trains the neural network with a given number of iterations by executing
// forward and back propagation
func (network *Network) Train(iterations int) {
	for i := 0; i < iterations; i++ {
		network.FeedForward()
		network.FeedBackward()
	}
}
