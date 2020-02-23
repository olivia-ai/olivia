package network

type Matrix [][]float64

type Network struct {
	Layers  []Matrix
	Weights [][][]float64
	Biases  [][][]float64
	Output  [][]float64
	Rate    float64
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(rate float64, input, output [][]float64, hiddensNodes ...int) Network {
	// Create the layers arrays and add the input values
	inputMatrix := input
	layers := []Matrix{inputMatrix}

	// Generate the hidden layers
	for _, hiddenNodes := range hiddensNodes {
		hiddenMatrix := CreateMatrix(len(input), hiddenNodes)
		layers = append(layers, hiddenMatrix)
	}

	// Add the output values to the layers arrays
	layers = append(layers, output)

	// Generate the weights and biases
	weightsNumber := 1 + len(hiddensNodes)
	var weights [][][]float64
	var biases [][][]float64

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
	z := Multiplication(
		ApplyFunction(Difference(output, network.Layers[2]), Twice),
		Multiplication(network.Layers[2], ApplyFunction(network.Layers[2], RemoveOne)),
	)
	w := DotProduct(Transpose(network.Layers[1]), z)
	network.Weights[1] = Sum(network.Weights[1], w)
	network.Biases[1] = Sum(network.Biases[1], z)

	z = Multiplication(
		DotProduct(
			z,
			Transpose(network.Weights[1]),
		),
		Multiplication(
			network.Layers[1],
			ApplyFunction(network.Layers[1], RemoveOne),
		),
	)

	network.Weights[0] = Sum(
		network.Weights[0],
		DotProduct(
			Transpose(network.Layers[0]),
			z,
		),
	)
	network.Biases[0] = Sum(network.Biases[0], z)
}
