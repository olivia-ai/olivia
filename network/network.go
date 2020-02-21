package network

type Network struct {
	layers       []Matrix
	zLayers      []Matrix
	weights      []Matrix
	biases       []Matrix
	actualOutput Matrix
}

// CreateNetwork creates the network by generating the layers, weights and biases
func CreateNetwork(input, output [][]float64, hiddensNodes ...int) Network {
	layers := []Matrix{
		{input},
	}
	zLayers := []Matrix{
		{input},
	}

	// Generate the hidden layers
	for _, hiddenNodes := range hiddensNodes {
		layers = append(layers, CreateMatrix(len(input), hiddenNodes))
		zLayers = append(zLayers, CreateMatrix(len(input), hiddenNodes))
	}

	layers = append(layers, Matrix{output})
	zLayers = append(layers, Matrix{output})

	// Generate the weights
	weightsNumber := 1 + len(hiddensNodes)
	var weights []Matrix
	var biases []Matrix

	for i := 0; i < weightsNumber; i++ {
		rows, columns := layers[i].Columns(), layers[i+1].Columns()
		weights = append(weights, RandomMatrix(rows, columns))

		biases = append(biases, RandomMatrix(layers[i].Rows(), columns))
	}

	return Network{
		layers:       layers,
		zLayers:      layers,
		actualOutput: Matrix{output},
		weights:      weights,
		biases:       biases,
	}
}

func (network *Network) FeedForward() {
	for i := 0; i < len(network.layers)-1; i++ {
		layer, weights, biases := network.layers[i], network.weights[i], network.biases[i]
		productMatrix := layer.DotProduct(weights)
		productMatrix.Add(biases)
		network.zLayers[i+1] = productMatrix
		productMatrix.ApplyFunction(Sigmoid)

		network.layers[i+1] = productMatrix
	}
}

func (network *Network) FeedBackward() {
	input := network.layers[len(network.layers)-1]
	lastInput := network.layers[len(network.layers)-2]
	error := network.actualOutput.Remove(input)
	z := network.zLayers[len(network.zLayers)-1]

	error = error.ApplyFunction(Twice)
	z = z.ApplyFunction(SigmoidDerivative)
	lastInput.Transpose()

	product := lastInput.DotProduct(error.Multiply(z))

	network.weights[len(network.weights)-1] = network.weights[len(network.weights)-1].Add(product)

	Test(network)
}

func Test(network *Network) {
	input := network.layers[len(network.layers)-1]
	error := network.actualOutput.Remove(input)
	z := network.zLayers[len(network.zLayers)-1]

	error = error.ApplyFunction(Twice)
	z = z.ApplyFunction(SigmoidDerivative)

	m := error.Multiply(z)
	m.Transpose()

	product := network.weights[len(network.weights)-1].DotProduct(m)

	network.weights[len(network.weights)-1] = network.weights[len(network.weights)-1].Add(product)
}
