package network

type Layer struct {
	Inputs  Matrix
	Outputs Matrix
}

type Network struct {
	Layers  []Layer
	Weights []Matrix
	Biases  []Matrix
	Output  Matrix
	Rate    float64
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(rate float64, input, output [][]float64, hiddensNodes ...int) Network {
	// Create the layers arrays and add the input values
	inputMatrix := Matrix{input}
	layers := []Layer{
		{
			inputMatrix,
			inputMatrix,
		},
	}

	// Generate the hidden layers
	for _, hiddenNodes := range hiddensNodes {
		hiddenMatrix := CreateMatrix(len(input), hiddenNodes)
		hiddenLayer := Layer{
			hiddenMatrix,
			hiddenMatrix,
		}

		layers = append(layers, hiddenLayer)
	}

	// Add the output values to the layers arrays
	outputMatrix := Matrix{output}
	layers = append(layers, Layer{
		outputMatrix,
		outputMatrix,
	})

	// Generate the weights and biases
	weightsNumber := 1 + len(hiddensNodes)
	var weights []Matrix
	var biases []Matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := layers[i].Inputs.Columns(), layers[i+1].Inputs.Columns()

		weights = append(weights, RandomMatrix(rows, columns))
		biases = append(biases, RandomMatrix(layers[i].Inputs.Rows(), columns))
	}

	return Network{
		Layers:  layers,
		Weights: weights,
		Biases:  biases,
		Output:  Matrix{output},
		Rate:    rate,
	}
}

// FeedForward executes forward propagation for the given inputs in the network
func (network *Network) FeedForward() {
	for i := 0; i < len(network.Layers)-1; i++ {
		layer, weights, biases := network.Layers[i].Outputs, network.Weights[i], network.Biases[i]

		productMatrix := layer.DotProduct(weights)
		productMatrix.Add(biases)

		// Replace the input values
		network.Layers[i+1].Inputs = productMatrix
		productMatrix.ApplyFunction(Sigmoid)

		// Replace the output values
		network.Layers[i+1].Outputs = productMatrix
	}
}

// FeedBackward executes back propagation to adjust the weights for all the layers
func (network *Network) FeedBackward() {
	output := Matrix{
		[][]float64{
			{1},
			{1},
			{0},
			{1},
			{1},
			{0},
		},
	}

	z := output.Substract(network.Layers[2].Outputs).ApplyFunction(Twice).Multiply(
		network.Layers[2].Outputs.Multiply(network.Layers[2].Outputs.ApplyFunction(RemoveOne)),
	)
	w := network.Layers[1].Outputs.Transpose().DotProduct(z)
	network.Weights[1] = network.Weights[1].Add(w)

	network.Weights[0] = network.Weights[0].Add(
		z.DotProduct(
			network.Weights[1].Transpose(),
		).Multiply(network.Layers[1].Outputs.Multiply(network.Layers[1].Outputs.ApplyFunction(RemoveOne))),
	)
}
